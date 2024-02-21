// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package statsagent

import (
	"context"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes"

	pb "github.com/DataDog/datadog-agent/pkg/proto/pbgo/trace"
	"github.com/DataDog/datadog-agent/pkg/trace/agent"
	"github.com/DataDog/datadog-agent/pkg/trace/api"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/metrics"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"

	traceconfig "github.com/DataDog/datadog-agent/pkg/trace/config"
	"github.com/DataDog/datadog-agent/pkg/trace/stats"
	"github.com/DataDog/datadog-agent/pkg/trace/telemetry"
	"github.com/DataDog/datadog-agent/pkg/trace/timing"
)

type StatsAgent interface {
	Start()
	Stop()
	ComputeStats(ctx context.Context, traces ptrace.Traces)
	ComputeDDStats(ctx context.Context, payload *api.Payload)
}

type StatsAgentConfig struct {
	ComputeStatsBySpanKind bool
	PeerTagsAggregation    bool
	SpanNameAsResourceName bool
	SpanNameRemappings     map[string]string
	IgnoreResources        []string
	PeerTags               []string
}

type statsAgent struct {
	out chan *pb.StatsPayload
	*agent.Agent

	// pchan specifies the channel that will be used to output Datadog Trace Agent API Payloads
	// resulting from ingested OpenTelemetry spans.
	pchan chan *api.Payload

	// wg waits for all goroutines to exit.
	wg sync.WaitGroup

	// senderWG waits for all sender goroutines to exit.
	senderWG sync.WaitGroup

	// exit signals the agent to shut down.
	exit chan struct{}

	// stopWriting signals to stop sending payloads to the agent
	stopWriting chan struct{}
}

func New(ctx context.Context, cfg *StatsAgentConfig, out chan *pb.StatsPayload, statsd statsd.ClientInterface) (StatsAgent, error) {
	acfg := traceconfig.New()
	acfg.OTLPReceiver.SpanNameRemappings = cfg.SpanNameRemappings
	acfg.OTLPReceiver.SpanNameAsResourceName = cfg.SpanNameAsResourceName
	acfg.Ignore["resource"] = cfg.IgnoreResources
	acfg.ComputeStatsBySpanKind = cfg.ComputeStatsBySpanKind
	acfg.PeerTagsAggregation = cfg.PeerTagsAggregation
	acfg.PeerTags = cfg.PeerTags

	// disable the HTTP receiver
	acfg.ReceiverPort = 0
	// set the API key to succeed startup; it is never used nor needed
	acfg.Endpoints[0].APIKey = "skip_check"
	// set the default hostname to the translator's placeholder; in the case where no hostname
	// can be deduced from incoming traces, we don't know the default hostname (because it is set
	// in the exporter). In order to avoid duplicating the hostname setting in the processor and
	// exporter, we use a placeholder and fill it in later (in the Datadog Exporter or Agent OTLP
	// Ingest). This gives a better user experience.
	acfg.Hostname = metrics.UnsetHostnamePlaceholder
	comonentSettings := component.TelemetrySettings{
		// TODO  : Need to update the below settings
		Logger:         zap.NewNop(),
		TracerProvider: nooptrace.NewTracerProvider(),
		MeterProvider:  noopmetric.NewMeterProvider(),
		MetricsLevel:   configtelemetry.LevelNone,
		Resource:       pcommon.NewResource(),
		ReportStatus: func(*component.StatusEvent) {
		},
	}

	attributesTranslator, err := attributes.NewTranslator(comonentSettings)
	if err != nil {
		return nil, err
	}
	acfg.OTLPReceiver.AttributesTranslator = attributesTranslator
	pchan := make(chan *api.Payload, 1000)
	a := agent.NewAgent(ctx, acfg, telemetry.NewNoopCollector(), statsd)
	// replace the Concentrator (the component which computes and flushes APM Stats from incoming
	// traces) with our own, which uses the 'out' channel.
	a.Concentrator = stats.NewConcentrator(acfg, out, time.Now(), statsd)
	// ...and the same for the ClientStatsAggregator; we don't use it here, but it is also a source
	// of stats which should be available to us.
	// a.ClientStatsAggregator = stats.NewClientStatsAggregator(acfg, out, statsd)
	// lastly, start the OTLP receiver, which will be used to introduce ResourceSpans into the traceagent,
	// so that we can transform them to Datadog spans and receive stats.
	timing := timing.New(statsd)

	a.OTLPReceiver = api.NewOTLPReceiver(pchan, acfg, statsd, timing)
	return &statsAgent{
		Agent: a,
		out:   out,
		pchan: pchan,
		exit:  make(chan struct{}),
	}, nil
}

// Start starts the traceagent, making it ready to ingest spans.
func (p *statsAgent) Start() {
	// we don't need to start the full agent, so we only start a set of minimal
	// components needed to compute stats:
	for _, starter := range []interface{ Start() }{
		p.Concentrator,
		// p.ClientStatsAggregator,
		// we don't need the samplers' nor the processor's functionalities;
		// but they are used by the agent nevertheless, so they need to be
		// active and functioning.
		p.PrioritySampler,
		p.ErrorsSampler,
		p.NoPrioritySampler,
		p.EventProcessor,
	} {
		starter.Start()
	}

	p.goDrain()
	p.goProcess()
}

// Stop stops the traceagent, making it unable to ingest spans. Do not call Ingest after Stop.
func (p *statsAgent) Stop() {
	for _, stopper := range []interface{ Stop() }{
		p.Concentrator,
		// p.ClientStatsAggregator,
		p.PrioritySampler,
		p.ErrorsSampler,
		p.NoPrioritySampler,
		p.EventProcessor,
	} {
		stopper.Stop()
	}
	close(p.exit)
	p.senderWG.Wait() //We need to stop sending payloads before we can close the traceWriter.In channel
	close(p.TraceWriter.In)
	p.wg.Wait()
}

// goDrain drains the TraceWriter channel, ensuring it won't block. We don't need the traces,
// nor do we have a running TraceWrite. We just want the outgoing stats.
func (p *statsAgent) goDrain() {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		for {
			select {
			case _, isOpen := <-p.TraceWriter.In:
				if !isOpen {
					return
				}
			}
		}
	}()
}

// ComputeStats processes the given spans within the traceagent and outputs stats through the output channel
// provided to newAgent. Do not call Ingest on an unstarted or stopped traceagent.
func (p *statsAgent) ComputeStats(ctx context.Context, traces ptrace.Traces) {
	rspanss := traces.ResourceSpans()
	for i := 0; i < rspanss.Len(); i++ {
		rspans := rspanss.At(i)
		p.OTLPReceiver.ReceiveResourceSpans(ctx, rspans, http.Header{})
		// ...the call transforms the OTLP Spans into a Datadog payload and sends the result
		// down the p.pchan channel

	}
}

func (p *statsAgent) ComputeDDStats(ctx context.Context, payload *api.Payload) {
	select {
	case p.pchan <- payload:
		return
	case <-ctx.Done():
		return
	}
}

// goProcesses runs the main loop which takes incoming payloads, processes them and generates stats.
// It then picks up those stats and converts them to metrics.
func (p *statsAgent) goProcess() {
	for i := 0; i < runtime.NumCPU(); i++ {
		p.senderWG.Add(1)
		go func() {
			defer p.senderWG.Done()
			for {
				select {
				case payload := <-p.pchan:
					p.Process(payload)
					// ...the call processes the payload and outputs stats via the 'out' channel
					// provided to newAgent
				case <-p.exit:
					return
				}
			}
		}()
	}
}

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

//go:build test

package npschedulerimpl

import (
	"fmt"
	"testing"
	"time"

	model "github.com/DataDog/agent-payload/v5/process"
	"github.com/DataDog/datadog-agent/comp/aggregator/demultiplexer/demultiplexerimpl"
	"github.com/DataDog/datadog-agent/comp/core"
	"github.com/DataDog/datadog-agent/comp/core/sysprobeconfig/sysprobeconfigimpl"
	"github.com/DataDog/datadog-agent/comp/forwarder/defaultforwarder"
	"github.com/DataDog/datadog-agent/comp/forwarder/eventplatform/eventplatformimpl"
	"github.com/DataDog/datadog-agent/comp/ndmtmp/forwarder/forwarderimpl"
	"github.com/DataDog/datadog-agent/comp/networkpath/npscheduler"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

// MockTimeNow mocks time.Now
var MockTimeNow = func() time.Time {
	layout := "2006-01-02 15:04:05"
	str := "2000-01-01 00:00:00"
	t, _ := time.Parse(layout, str)
	return t
}

// testOptions is a fx collection of common dependencies for all tests
var testOptions = fx.Options(
	Module(),
	forwarderimpl.MockModule(),
	demultiplexerimpl.MockModule(),
	defaultforwarder.MockModule(),
	core.MockBundle(),
	eventplatformimpl.MockModule(),
)

func newTestNpScheduler(t *testing.T, sysConfigs map[string]any) (*fxtest.App, *npSchedulerImpl) {
	var component npscheduler.Component
	app := fxtest.New(t, fx.Options(
		testOptions,
		fx.Supply(fx.Annotate(t, fx.As(new(testing.TB)))),
		fx.Replace(sysprobeconfigimpl.MockParams{Overrides: sysConfigs}),
		fx.Populate(&component),
	))
	npScheduler := component.(*npSchedulerImpl)

	require.NotNil(t, npScheduler)
	require.NotNil(t, app)
	return app, npScheduler
}

func createConns(numberOfConns int) []*model.Connection {
	var conns []*model.Connection
	for i := 0; i < numberOfConns; i++ {
		conns = append(conns, &model.Connection{
			Laddr:     &model.Addr{Ip: fmt.Sprintf("127.0.0.%d", i), Port: int32(30000)},
			Raddr:     &model.Addr{Ip: fmt.Sprintf("127.0.1.%d", i), Port: int32(80)},
			Direction: model.ConnectionDirection_outgoing,
		})
	}
	return conns
}

func waitForProcessedPathtests(npScheduler *npSchedulerImpl, timeout time.Duration, processecCount uint64) {
	timeoutChan := time.After(timeout)
	tick := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-timeoutChan:
			return
		case <-tick:
			if npScheduler.processedTracerouteCount.Load() >= processecCount {
				return
			}
		}
	}
}

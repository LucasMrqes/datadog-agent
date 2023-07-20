// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux

package modules

import (
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"time"

	"go.uber.org/atomic"

	"github.com/DataDog/datadog-agent/cmd/system-probe/api/module"
	"github.com/DataDog/datadog-agent/cmd/system-probe/config"
	"github.com/DataDog/datadog-agent/cmd/system-probe/utils"
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/ebpf/probe/ebpfcheck"
	"github.com/DataDog/datadog-agent/pkg/ebpf"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// EBPFProbe Factory
var EBPFProbe = module.Factory{
	Name:             config.EBPFModule,
	ConfigNamespaces: []string{},
	Fn: func(cfg *config.Config) (module.Module, error) {
		log.Infof("Starting the ebpf probe")
		okp, err := ebpfcheck.NewEBPFProbe(ebpf.NewConfig())
		if err != nil {
			return nil, fmt.Errorf("unable to start the ebpf probe: %w", err)
		}
		return &ebpfModule{
			EBPFProbe: okp,
			lastCheck: atomic.NewInt64(0),
		}, nil
	},
}

var _ module.Module = &ebpfModule{}

type ebpfModule struct {
	*ebpfcheck.EBPFProbe
	lastCheck *atomic.Int64
}

func (o *ebpfModule) RegisterGRPC(server *grpc.Server) error {
	return nil
}

func (o *ebpfModule) Register(httpMux *module.Router) error {
	httpMux.HandleFunc("/check", utils.WithConcurrencyLimit(utils.DefaultMaxConcurrentRequests, func(w http.ResponseWriter, req *http.Request) {
		o.lastCheck.Store(time.Now().Unix())
		stats := o.EBPFProbe.GetAndFlush()
		utils.WriteAsJSON(w, stats)
	}))

	return nil
}

func (o *ebpfModule) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"last_check": o.lastCheck.Load(),
	}
}

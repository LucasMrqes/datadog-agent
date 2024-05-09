// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

package npschedulerimpl

import (
	"context"

	"github.com/DataDog/datadog-agent/comp/core/log"
	"github.com/DataDog/datadog-agent/comp/core/sysprobeconfig"
	"github.com/DataDog/datadog-agent/comp/forwarder/eventplatform"
	"github.com/DataDog/datadog-agent/comp/networkpath/npscheduler"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
	"go.uber.org/fx"
)

type dependencies struct {
	fx.In
	Lc          fx.Lifecycle
	EpForwarder eventplatform.Component
	Logger      log.Component
	Sysconfig   sysprobeconfig.Component
}

type provides struct {
	fx.Out

	Comp npscheduler.Component
}

// Module defines the fx options for this component.
func Module() fxutil.Module {
	return fxutil.Component(
		fx.Provide(newNpScheduler),
	)
}

func newNpScheduler(deps dependencies) provides {
	var scheduler *npSchedulerImpl

	networkPathEnabled := deps.Sysconfig.GetBool("network_path.enabled")
	if networkPathEnabled {
		deps.Logger.Debugf("Network Path Scheduler enabled")
		scheduler = newNpSchedulerImpl(deps.EpForwarder, deps.Logger, deps.Sysconfig)
		deps.Lc.Append(fx.Hook{
			// No need for OnStart hook since NpScheduler.Init() will be called by clients when needed.
			OnStart: func(context.Context) error {
				scheduler.start()
				return nil
			},
			OnStop: func(context.Context) error {
				scheduler.stop()
				return nil
			},
		})
	} else {
		deps.Logger.Debugf("Network Path Scheduler disabled")
		scheduler = newNoopNpSchedulerImpl()
	}

	return provides{
		Comp: scheduler,
	}
}
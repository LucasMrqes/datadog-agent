// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package fxrunner

import (
	"github.com/DataDog/datadog-agent/comp/metadata/runner/impl"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
)

// Module specifies the compression module.
func Module() fxutil.Module {
	return fxutil.Component(
		fxutil.ProvideComponentConstructor(
			impl.NewRunner,
		),
	)
}

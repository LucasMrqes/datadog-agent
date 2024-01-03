// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !windows

// Package wincrashdetect implements the windows crash detection on windows.  It does nothing on linux
package wincrashdetect

import "github.com/DataDog/datadog-agent/pkg/collector/check"

const (
	Enabled   = false
	CheckName = "wincrashdetect"
)

func Factory() check.Check {
	return nil
}

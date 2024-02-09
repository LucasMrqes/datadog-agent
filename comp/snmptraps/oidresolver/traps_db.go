// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022-present Datadog, Inc.

package oidresolver

import (
	"strings"
	"unicode"
)

// VariableMetadata is the MIB-extracted information of a given trap variable
type VariableMetadata struct {
	Name               string         `yaml:"name" json:"name"`
	Description        string         `yaml:"descr" json:"descr"`
	Enumeration        map[int]string `yaml:"enum" json:"enum"`
	Bits               map[int]string `yaml:"bits" json:"bits"`
	IsIntermediateNode bool           `yaml:"-" json:"-"`
	// In theory, variables should always be leaves of the OID tree as intermediate nodes do not contain data.
	// This isn't true in practice (see 1.3.6.1.4.1.4962.2.1.6.3).
	// Variables are resolved by 'climbing' up the OID tree until finding a match, but variables that are known to be nodes
	// should never be used for resolving.
}

// VariableSpec contains the variableMetadata for each known variable of a given trap db file
type VariableSpec map[string]VariableMetadata

// TrapMetadata is the MIB-extracted information of a given trap OID.
// It also contains a reference to the variableSpec that was defined in the same trap db file.
// This is to prevent variable conflicts and to give precedence to the variable definitions located]
// in the same trap db file as the trap.
type TrapMetadata struct {
	Name            string       `yaml:"name" json:"name"`
	MIBName         string       `yaml:"mib" json:"mib"`
	Description     string       `yaml:"descr" json:"descr"`
	VariableSpecPtr VariableSpec `yaml:"-" json:"-"`
}

// TrapSpec contains the variableMetadata for each known trap in all trap db files
type TrapSpec map[string]TrapMetadata

// TrapDBFileContent contains data for the traps and variables from a trap db file.
type TrapDBFileContent struct {
	Traps     TrapSpec     `yaml:"traps" json:"traps"`
	Variables VariableSpec `yaml:"vars" json:"vars"`
}

// NormalizeOID converts an OID from the absolute form ".1.2.3..." to a relative form "1.2.3..."
func NormalizeOID(value string) string {
	// OIDs can be formatted as ".1.2.3..." ("absolute form") or "1.2.3..." ("relative form").
	// Convert everything to relative form, like we do in the Python check.
	return strings.TrimLeft(value, ".")
}

// IsValidOID returns true if value looks like a valid OID.
// An OID is made of digits and dots, but OIDs do not end with a dot and there are always
// digits between dots.
func IsValidOID(value string) bool {
	var previousChar rune
	for _, char := range value {
		if char != '.' && !unicode.IsDigit(char) {
			return false
		}
		if char == '.' && previousChar == '.' {
			return false
		}
		previousChar = char
	}
	return previousChar != '.'
}

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package iis

import (
	//"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IISSiteDefinition represents an IIS site definition
type IISSiteDefinition struct {
	Name        string //  name of the site
	BindingPort string // port to bind to, of the form '*:8081'
	AssetsDir   string // directory to copy for assets
}

// Configuration represents the Active Directory configuration (domain name, password, users etc...)
type Configuration struct {
	Sites []IISSiteDefinition
}

// Option is an optional function parameter type for Configuration options
type Option = func(*Configuration) error

// WithSite adds a new IIS site to the configuration
func WithSite(name, bindingport, assetsdir string) func(*Configuration) error {
	return func(p *Configuration) error {
		p.Sites = append(p.Sites, IISSiteDefinition{
			Name:        name,
			BindingPort: bindingport,
			AssetsDir:   assetsdir,
		})
		return nil
	}
}

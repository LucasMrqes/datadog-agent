// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

// Package rdnsquerierimpl provides JMW impl or not?
package rdnsquerierimpl

import (
	"fmt"
	"net"
	"time"

	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/comp/core/log"
	"github.com/DataDog/datadog-agent/comp/rdnsquerier"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
	//JMW"github.com/DataDog/datadog-agent/comp/core/config"
	//JMWNOTUSED nfconfig "github.com/DataDog/datadog-agent/comp/netflow/config"
)

type dependencies struct {
	fx.In
	Lc     fx.Lifecycle
	Logger log.Component
	// JMWTELEMETRY dependency?
}

type provides struct {
	fx.Out
	Comp rdnsquerier.Component
}

// Module defines the fx options for this component.
func Module() fxutil.Module {
	return fxutil.Component(
		fx.Provide(newRDNSQuerier),
	)
}

func newRDNSQuerier(deps dependencies) provides {
	// Component initialization
	rdnsQuerier := &rdnsQuerier{
		lc:     deps.Lc,
		logger: deps.Logger,
		cache:  make(map[string]rdnsCacheEntry),
	}
	return provides{
		Comp: rdnsQuerier,
	}
}

type rdnsCacheEntry struct {
	//JMWhostname string
	//JMWUNUSED expirationTime int64
	// map of hashes to callback to set hostname
	//JMWcallbacks map[string]func(string)
}

// RDNSQuerier provides JMW
type rdnsQuerier struct {
	lc     fx.Lifecycle
	logger log.Component

	// mutex for JMW
	//JMWmutex sync.RWMutex

	// map of ip to hostname and expiration time
	cache map[string]rdnsCacheEntry
}

/*JMWRM no longer needed now that it's a component, but still used in aggregator_test.go
// NewRDNSQuerier creates a new RDNSQuerier JMW component.
func NewRDNSQuerier() *RDNSQuerier {
	return &RDNSQuerier{
		cache: make(map[string]rdnsCacheEntry),
	}
}
*/

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s: took %v usec\n", name, time.Since(start).Microseconds())
	}
}

// GetHostname returns the hostname for the given IP address JMW
func (q *rdnsQuerier) GetHostname(ipAddr []byte) string {
	defer timer("timer JMW GetHostname() all")()

	ip := net.IP(ipAddr)
	if !ip.IsPrivate() { // JMW IsPrivate() also returns false for invalid IP addresses JMWCHECK
		fmt.Printf("JMW GetHostname() IP address `%s` is not private\n", ip.String())
		// JMWTELEMETRY increment NOT private IP address counter
		return ""
	}

	// JMWTELEMETRY increment private IP address counter
	addr := ip.String()
	// JMW LookupAddr can return both a non-zero length slice of hostnames and an error.
	// BUT When using the host C library resolver, at most one result will be returned.
	// So for now, when specifying DNS resolvers is not supported, if we get an error we know that there is no valid hostname returned.
	// If/when we add support for specifying DNS resolvers, there may be multiple hostnames returned, and there may be one or more hostname returned AT THE SAME TIME an error is returned.  To keep it simple, if there is no error, we will just return the first hostname, and if there is an error, we will return an empty string and add telemetry about the error.
	defer timer("timer JMW GetHostname() LookupAddr")()
	hostnames, err := net.LookupAddr(addr)
	if err != nil {
		//JMWADDLOGGER f.logger.Warnf("JMW Failed to lookup hostname for IP address `%s`: %s", addr, err)
		fmt.Printf("JMW GetHostname() error looking up hostname for IP address `%s`: %s\n", addr, err)
		// JMWTELEMETRY increment metric for failed lookups - JMW should I differentiate between no match and other errors? or just tag w/ error?  how to tag w/ error w/out the tag being a string (unlimited cardinality)?
		return ""
	}

	if len(hostnames) == 0 { // JMW is this even possible? // JMWRM?
		fmt.Printf("JMW IP address `%s` has no match - returning empty hostname", addr)
		// JMWTELEMETRY increment metric for no match
		return ""
	}

	// JMWTELEMETRY increment metric for successful lookups
	//if (len(hostnames) > 1) {
	// JMWTELEMETRY increment metric for multiple hostnames
	//}
	fmt.Printf("JMW GetHostname() IP address `%s` matched - returning hostname `%s`\n", addr, hostnames[0])
	return hostnames[0]
}

/*
// JMW Get returns the hostname for the given IP address
func (q *rdnsQuerier) Get(ip string) string {
	entry, ok := q.cache[ip]
	if ok && entry.expirationTime < time.Now().Unix() {
		return entry.hostname
	}

	return entry.hostname
}
*/

/* JMWASYNC
func (q *rdnsQuerier) GetAsync(ip string, func inlineCallback(string), func asyncCallback(string)) {
	entry, ok := q.cache[ip]
	if ok {
		if entry.expirationTime < time.Now().Unix() {
			inlineCallback(entry)
		}
		return
	}
	if entry.expirationTime < time.Now().Unix() {
		func()
		return
	}
	asyncCallback(entry.hostname)
}
*/

/*
type reverseDNSCache struct {
	// JMW IP address to hostname
	cache map[string]string

	// JMW mutex for cache
	mutex sync.RWMutex
}

func NewReverseDNSCache func() *reverseDNSCache {
	return &reverseDNSCache{
		cache: make(map[string]string),
	}
}

func (r *reverseDNSCache) PreFetch(ip string) string {
}
func (r *reverseDNSCache) Expire() string {
}
func (r *reverseDNSCache) TryGet(ip string) (string, bool) {
}
*/

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fqdn

import (
	"github.com/cilium/cilium/pkg/ipcache"
)

// Config is a simple configuration structure to set how pkg/fqdn subcomponents
// behave.
type Config struct {
	// MinTTL is the time used by the poller to cache information.
	MinTTL int

	// Cache is where the poller stores DNS data used to generate rules.
	// When set to nil, it uses fqdn.DefaultDNSCache, a global cache instance.
	Cache *DNSCache

	// GetEndpointsDNSInfo is a function that returns a list of fqdn-relevant fields from all Endpoints known to the agent.
	// The endpoint's DNSHistory and DNSZombies are used as part of the garbage collection and restoration processes.
	//
	// Optional parameter endpointID will cause this function to only return the endpoint with the specified ID.
	GetEndpointsDNSInfo func(endpointID string) []EndpointDNSInfo

	IPCache IPCache
}

type EndpointDNSInfo struct {
	ID         string
	ID64       int64
	DNSHistory *DNSCache
	DNSZombies *DNSZombieMappings
}

type IPCache interface {
	UpsertMetadataBatch(updates ...ipcache.MU) (revision uint64)
	RemoveMetadataBatch(updates ...ipcache.MU) (revision uint64)
	WaitForRevision(rev uint64)
}

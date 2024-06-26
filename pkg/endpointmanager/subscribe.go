// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package endpointmanager

import "github.com/cilium/cilium/pkg/endpoint"

// Subscribers may register via Subscribe() to be notified when events occur.
type Subscriber interface {
	// EndpointCreated is called at the end of endpoint creation.
	// Implementations must not attempt write operations on the
	// EndpointManager from this callback.
	// This function is being called inside a RLock, so it must not attempt
	// to acquire a lock on the EndpointManager.
	EndpointCreated(ep *endpoint.Endpoint)

	// EndpointDeleted is called at the end of endpoint deletion.
	// Implementations must not attempt write operations on the
	// EndpointManager from this callback.
	// This function is being called inside a RLock, so it must not attempt
	// to acquire a lock on the EndpointManager.
	EndpointDeleted(ep *endpoint.Endpoint, conf endpoint.DeleteConfig)

	// EndpointRestored is called at the end of endpoint restoration.
	// Implementations must not attempt write operations on the
	// EndpointManager from this callback.
	// This function is being called inside a RLock, so it must not attempt
	// to acquire a lock on the EndpointManager.
	EndpointRestored(ep *endpoint.Endpoint)
}

func (mgr *endpointManager) Subscribe(s Subscriber) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	mgr.subscribers[s] = struct{}{}
}

func (mgr *endpointManager) Unsubscribe(s Subscriber) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()
	delete(mgr.subscribers, s)
}

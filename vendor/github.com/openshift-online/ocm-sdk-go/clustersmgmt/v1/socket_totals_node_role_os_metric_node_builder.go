/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// SocketTotalsNodeRoleOSMetricNodeBuilder contains the data and logic needed to build 'socket_totals_node_role_OS_metric_node' objects.
//
// Representation of information from telemetry about the socket capacity by node
// role and OS of a cluster.
type SocketTotalsNodeRoleOSMetricNodeBuilder struct {
	socketTotals []*SocketTotalNodeRoleOSMetricNodeBuilder
}

// NewSocketTotalsNodeRoleOSMetricNode creates a new builder of 'socket_totals_node_role_OS_metric_node' objects.
func NewSocketTotalsNodeRoleOSMetricNode() *SocketTotalsNodeRoleOSMetricNodeBuilder {
	return new(SocketTotalsNodeRoleOSMetricNodeBuilder)
}

// SocketTotals sets the value of the 'socket_totals' attribute
// to the given values.
//
//
func (b *SocketTotalsNodeRoleOSMetricNodeBuilder) SocketTotals(values ...*SocketTotalNodeRoleOSMetricNodeBuilder) *SocketTotalsNodeRoleOSMetricNodeBuilder {
	b.socketTotals = make([]*SocketTotalNodeRoleOSMetricNodeBuilder, len(values))
	copy(b.socketTotals, values)
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *SocketTotalsNodeRoleOSMetricNodeBuilder) Copy(object *SocketTotalsNodeRoleOSMetricNode) *SocketTotalsNodeRoleOSMetricNodeBuilder {
	if object == nil {
		return b
	}
	if object.socketTotals != nil && len(object.socketTotals.items) > 0 {
		b.socketTotals = make([]*SocketTotalNodeRoleOSMetricNodeBuilder, len(object.socketTotals.items))
		for i, item := range object.socketTotals.items {
			b.socketTotals[i] = NewSocketTotalNodeRoleOSMetricNode().Copy(item)
		}
	} else {
		b.socketTotals = nil
	}
	return b
}

// Build creates a 'socket_totals_node_role_OS_metric_node' object using the configuration stored in the builder.
func (b *SocketTotalsNodeRoleOSMetricNodeBuilder) Build() (object *SocketTotalsNodeRoleOSMetricNode, err error) {
	object = new(SocketTotalsNodeRoleOSMetricNode)
	if b.socketTotals != nil {
		object.socketTotals = new(SocketTotalNodeRoleOSMetricNodeList)
		object.socketTotals.items = make([]*SocketTotalNodeRoleOSMetricNode, len(b.socketTotals))
		for i, item := range b.socketTotals {
			object.socketTotals.items[i], err = item.Build()
			if err != nil {
				return
			}
		}
	}
	return
}

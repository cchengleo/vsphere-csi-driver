/*
Copyright 2019 The Kubernetes Authors.

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

package vsphere

import (
	"context"

	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/vim25"
	"k8s.io/klog"
)

// NewCNSClient creates a new CNS client
func NewCNSClient(ctx context.Context, c *vim25.Client) (*cns.Client, error) {
	cnsClient, err := cns.NewClient(ctx, c)
	if err != nil {
		klog.Errorf("Failed to create a new client for CNS. err: %v", err)
		return nil, err
	}
	return cnsClient, nil
}

// ConnectCNS creates a CNS client for the virtual center.
func (vc *VirtualCenter) ConnectCNS(ctx context.Context) error {
	err := vc.Connect(ctx)
	if err != nil {
		klog.Errorf("Failed to connect to Virtual Center host %q with err: %v", vc.Config.Host, err)
		return err
	}
	if vc.CnsClient == nil {
		if vc.CnsClient, err = NewCNSClient(ctx, vc.Client.Client); err != nil {
			klog.Errorf("Failed to create CNS client on vCenter host %q with err: %v", vc.Config.Host, err)
			return err
		}
	}
	return nil
}

// DisconnectCNS destroys the CNS client for the virtual center.
func (vc *VirtualCenter) DisconnectCNS(ctx context.Context) {
	if vc.CnsClient == nil {
		klog.V(1).Info("CnsClient wasn't connected, ignoring")
	} else {
		vc.CnsClient = nil
	}
}

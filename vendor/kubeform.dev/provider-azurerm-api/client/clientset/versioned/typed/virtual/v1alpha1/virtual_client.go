/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"
	"kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type VirtualV1alpha1Interface interface {
	RESTClient() rest.Interface
	DesktopApplicationsGetter
	DesktopApplicationGroupsGetter
	DesktopHostPoolsGetter
	DesktopWorkspacesGetter
	DesktopWorkspaceApplicationGroupAssociationsGetter
	HubsGetter
	HubBGPConnectionsGetter
	HubConnectionsGetter
	HubIPsGetter
	HubRouteTablesGetter
	HubSecurityPartnerProvidersGetter
	MachinesGetter
	MachineConfigurationPolicyAssignmentsGetter
	MachineDataDiskAttachmentsGetter
	MachineExtensionsGetter
	MachineScaleSetsGetter
	MachineScaleSetExtensionsGetter
	NetworksGetter
	NetworkGatewaysGetter
	NetworkGatewayConnectionsGetter
	NetworkPeeringsGetter
	WansGetter
}

// VirtualV1alpha1Client is used to interact with features provided by the virtual.azurerm.kubeform.com group.
type VirtualV1alpha1Client struct {
	restClient rest.Interface
}

func (c *VirtualV1alpha1Client) DesktopApplications(namespace string) DesktopApplicationInterface {
	return newDesktopApplications(c, namespace)
}

func (c *VirtualV1alpha1Client) DesktopApplicationGroups(namespace string) DesktopApplicationGroupInterface {
	return newDesktopApplicationGroups(c, namespace)
}

func (c *VirtualV1alpha1Client) DesktopHostPools(namespace string) DesktopHostPoolInterface {
	return newDesktopHostPools(c, namespace)
}

func (c *VirtualV1alpha1Client) DesktopWorkspaces(namespace string) DesktopWorkspaceInterface {
	return newDesktopWorkspaces(c, namespace)
}

func (c *VirtualV1alpha1Client) DesktopWorkspaceApplicationGroupAssociations(namespace string) DesktopWorkspaceApplicationGroupAssociationInterface {
	return newDesktopWorkspaceApplicationGroupAssociations(c, namespace)
}

func (c *VirtualV1alpha1Client) Hubs(namespace string) HubInterface {
	return newHubs(c, namespace)
}

func (c *VirtualV1alpha1Client) HubBGPConnections(namespace string) HubBGPConnectionInterface {
	return newHubBGPConnections(c, namespace)
}

func (c *VirtualV1alpha1Client) HubConnections(namespace string) HubConnectionInterface {
	return newHubConnections(c, namespace)
}

func (c *VirtualV1alpha1Client) HubIPs(namespace string) HubIPInterface {
	return newHubIPs(c, namespace)
}

func (c *VirtualV1alpha1Client) HubRouteTables(namespace string) HubRouteTableInterface {
	return newHubRouteTables(c, namespace)
}

func (c *VirtualV1alpha1Client) HubSecurityPartnerProviders(namespace string) HubSecurityPartnerProviderInterface {
	return newHubSecurityPartnerProviders(c, namespace)
}

func (c *VirtualV1alpha1Client) Machines(namespace string) MachineInterface {
	return newMachines(c, namespace)
}

func (c *VirtualV1alpha1Client) MachineConfigurationPolicyAssignments(namespace string) MachineConfigurationPolicyAssignmentInterface {
	return newMachineConfigurationPolicyAssignments(c, namespace)
}

func (c *VirtualV1alpha1Client) MachineDataDiskAttachments(namespace string) MachineDataDiskAttachmentInterface {
	return newMachineDataDiskAttachments(c, namespace)
}

func (c *VirtualV1alpha1Client) MachineExtensions(namespace string) MachineExtensionInterface {
	return newMachineExtensions(c, namespace)
}

func (c *VirtualV1alpha1Client) MachineScaleSets(namespace string) MachineScaleSetInterface {
	return newMachineScaleSets(c, namespace)
}

func (c *VirtualV1alpha1Client) MachineScaleSetExtensions(namespace string) MachineScaleSetExtensionInterface {
	return newMachineScaleSetExtensions(c, namespace)
}

func (c *VirtualV1alpha1Client) Networks(namespace string) NetworkInterface {
	return newNetworks(c, namespace)
}

func (c *VirtualV1alpha1Client) NetworkGateways(namespace string) NetworkGatewayInterface {
	return newNetworkGateways(c, namespace)
}

func (c *VirtualV1alpha1Client) NetworkGatewayConnections(namespace string) NetworkGatewayConnectionInterface {
	return newNetworkGatewayConnections(c, namespace)
}

func (c *VirtualV1alpha1Client) NetworkPeerings(namespace string) NetworkPeeringInterface {
	return newNetworkPeerings(c, namespace)
}

func (c *VirtualV1alpha1Client) Wans(namespace string) WanInterface {
	return newWans(c, namespace)
}

// NewForConfig creates a new VirtualV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*VirtualV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &VirtualV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new VirtualV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *VirtualV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new VirtualV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *VirtualV1alpha1Client {
	return &VirtualV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *VirtualV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

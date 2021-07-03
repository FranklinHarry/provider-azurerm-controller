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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ServicesServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicesServerSpec   `json:"spec,omitempty"`
	Status            ServicesServerStatus `json:"status,omitempty"`
}

type ServicesServerSpecIpv4FirewallRule struct {
	Name       *string `json:"name" tf:"name"`
	RangeEnd   *string `json:"rangeEnd" tf:"range_end"`
	RangeStart *string `json:"rangeStart" tf:"range_start"`
}

type ServicesServerSpec struct {
	KubeformOutput *ServicesServerSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ServicesServerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ServicesServerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdminUsers []string `json:"adminUsers,omitempty" tf:"admin_users"`
	// +optional
	BackupBlobContainerURI *string `json:"-" sensitive:"true" tf:"backup_blob_container_uri"`
	// +optional
	EnablePowerBiService *bool `json:"enablePowerBiService,omitempty" tf:"enable_power_bi_service"`
	// +optional
	Ipv4FirewallRule []ServicesServerSpecIpv4FirewallRule `json:"ipv4FirewallRule,omitempty" tf:"ipv4_firewall_rule"`
	Location         *string                              `json:"location" tf:"location"`
	Name             *string                              `json:"name" tf:"name"`
	// +optional
	QuerypoolConnectionMode *string `json:"querypoolConnectionMode,omitempty" tf:"querypool_connection_mode"`
	ResourceGroupName       *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ServerFullName *string `json:"serverFullName,omitempty" tf:"server_full_name"`
	Sku            *string `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ServicesServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicesServerList is a list of ServicesServers
type ServicesServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicesServer CRD objects
	Items []ServicesServer `json:"items,omitempty"`
}

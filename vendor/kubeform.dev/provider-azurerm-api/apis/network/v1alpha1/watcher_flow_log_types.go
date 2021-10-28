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

type WatcherFlowLog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WatcherFlowLogSpec   `json:"spec,omitempty"`
	Status            WatcherFlowLogStatus `json:"status,omitempty"`
}

type WatcherFlowLogSpecRetentionPolicy struct {
	Days    *int64 `json:"days" tf:"days"`
	Enabled *bool  `json:"enabled" tf:"enabled"`
}

type WatcherFlowLogSpecTrafficAnalytics struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	IntervalInMinutes   *int64  `json:"intervalInMinutes,omitempty" tf:"interval_in_minutes"`
	WorkspaceID         *string `json:"workspaceID" tf:"workspace_id"`
	WorkspaceRegion     *string `json:"workspaceRegion" tf:"workspace_region"`
	WorkspaceResourceID *string `json:"workspaceResourceID" tf:"workspace_resource_id"`
}

type WatcherFlowLogSpec struct {
	State *WatcherFlowLogSpecResource `json:"state,omitempty" tf:"-"`

	Resource WatcherFlowLogSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WatcherFlowLogSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	Location               *string                            `json:"location,omitempty" tf:"location"`
	NetworkSecurityGroupID *string                            `json:"networkSecurityGroupID" tf:"network_security_group_id"`
	NetworkWatcherName     *string                            `json:"networkWatcherName" tf:"network_watcher_name"`
	ResourceGroupName      *string                            `json:"resourceGroupName" tf:"resource_group_name"`
	RetentionPolicy        *WatcherFlowLogSpecRetentionPolicy `json:"retentionPolicy" tf:"retention_policy"`
	StorageAccountID       *string                            `json:"storageAccountID" tf:"storage_account_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TrafficAnalytics *WatcherFlowLogSpecTrafficAnalytics `json:"trafficAnalytics,omitempty" tf:"traffic_analytics"`
	// +optional
	Version *int64 `json:"version,omitempty" tf:"version"`
}

type WatcherFlowLogStatus struct {
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

// WatcherFlowLogList is a list of WatcherFlowLogs
type WatcherFlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WatcherFlowLog CRD objects
	Items []WatcherFlowLog `json:"items,omitempty"`
}

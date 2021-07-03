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

type FlexibleServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServerSpec   `json:"spec,omitempty"`
	Status            FlexibleServerStatus `json:"status,omitempty"`
}

type FlexibleServerSpecMaintenanceWindow struct {
	// +optional
	DayOfWeek *int64 `json:"dayOfWeek,omitempty" tf:"day_of_week"`
	// +optional
	StartHour *int64 `json:"startHour,omitempty" tf:"start_hour"`
	// +optional
	StartMinute *int64 `json:"startMinute,omitempty" tf:"start_minute"`
}

type FlexibleServerSpec struct {
	KubeformOutput *FlexibleServerSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource FlexibleServerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type FlexibleServerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login"`
	// +optional
	AdministratorPassword *string `json:"-" sensitive:"true" tf:"administrator_password"`
	// +optional
	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`
	// +optional
	CmkEnabled *string `json:"cmkEnabled,omitempty" tf:"cmk_enabled"`
	// +optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode"`
	// +optional
	DelegatedSubnetID *string `json:"delegatedSubnetID,omitempty" tf:"delegated_subnet_id"`
	// +optional
	Fqdn     *string `json:"fqdn,omitempty" tf:"fqdn"`
	Location *string `json:"location" tf:"location"`
	// +optional
	MaintenanceWindow *FlexibleServerSpecMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	Name              *string                              `json:"name" tf:"name"`
	// +optional
	PointInTimeRestoreTimeInUtc *string `json:"pointInTimeRestoreTimeInUtc,omitempty" tf:"point_in_time_restore_time_in_utc"`
	// +optional
	PublicNetworkAccessEnabled *bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	ResourceGroupName          *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name"`
	// +optional
	SourceServerID *string `json:"sourceServerID,omitempty" tf:"source_server_id"`
	// +optional
	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type FlexibleServerStatus struct {
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

// FlexibleServerList is a list of FlexibleServers
type FlexibleServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FlexibleServer CRD objects
	Items []FlexibleServer `json:"items,omitempty"`
}

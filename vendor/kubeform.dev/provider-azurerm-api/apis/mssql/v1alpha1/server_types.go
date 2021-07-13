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

type Server struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec,omitempty"`
	Status            ServerStatus `json:"status,omitempty"`
}

type ServerSpecAzureadAdministrator struct {
	LoginUsername *string `json:"loginUsername" tf:"login_username"`
	ObjectID      *string `json:"objectID" tf:"object_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
}

type ServerSpecExtendedAuditingPolicy struct {
	// +optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled"`
	// +optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days"`
	// +optional
	StorageAccountAccessKey *string `json:"-" sensitive:"true" tf:"storage_account_access_key"`
	// +optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary"`
	// +optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint"`
}

type ServerSpecIdentity struct {
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type ServerSpec struct {
	State *ServerSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ServerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AdministratorLogin         *string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorLoginPassword *string `json:"-" sensitive:"true" tf:"administrator_login_password"`
	// +optional
	AzureadAdministrator *ServerSpecAzureadAdministrator `json:"azureadAdministrator,omitempty" tf:"azuread_administrator"`
	// +optional
	ConnectionPolicy *string `json:"connectionPolicy,omitempty" tf:"connection_policy"`
	// +optional
	// Deprecated
	ExtendedAuditingPolicy *ServerSpecExtendedAuditingPolicy `json:"extendedAuditingPolicy,omitempty" tf:"extended_auditing_policy"`
	// +optional
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty" tf:"fully_qualified_domain_name"`
	// +optional
	Identity *ServerSpecIdentity `json:"identity,omitempty" tf:"identity"`
	Location *string             `json:"location" tf:"location"`
	// +optional
	MinimumTlsVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version"`
	Name              *string `json:"name" tf:"name"`
	// +optional
	PublicNetworkAccessEnabled *bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	ResourceGroupName          *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RestorableDroppedDatabaseIDS []string `json:"restorableDroppedDatabaseIDS,omitempty" tf:"restorable_dropped_database_ids"`
	// +optional
	Tags    *map[string]string `json:"tags,omitempty" tf:"tags"`
	Version *string            `json:"version" tf:"version"`
}

type ServerStatus struct {
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

// ServerList is a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Server CRD objects
	Items []Server `json:"items,omitempty"`
}

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

type InteractiveQueryCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InteractiveQueryClusterSpec   `json:"spec,omitempty"`
	Status            InteractiveQueryClusterStatus `json:"status,omitempty"`
}

type InteractiveQueryClusterSpecComponentVersion struct {
	InteractiveHive *string `json:"interactiveHive" tf:"interactive_hive"`
}

type InteractiveQueryClusterSpecGateway struct {
	// +optional
	// Deprecated
	Enabled  *bool   `json:"enabled,omitempty" tf:"enabled"`
	Password *string `json:"-" sensitive:"true" tf:"password"`
	Username *string `json:"username" tf:"username"`
}

type InteractiveQueryClusterSpecMetastoresAmbari struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type InteractiveQueryClusterSpecMetastoresHive struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type InteractiveQueryClusterSpecMetastoresOozie struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type InteractiveQueryClusterSpecMetastores struct {
	// +optional
	Ambari *InteractiveQueryClusterSpecMetastoresAmbari `json:"ambari,omitempty" tf:"ambari"`
	// +optional
	Hive *InteractiveQueryClusterSpecMetastoresHive `json:"hive,omitempty" tf:"hive"`
	// +optional
	Oozie *InteractiveQueryClusterSpecMetastoresOozie `json:"oozie,omitempty" tf:"oozie"`
}

type InteractiveQueryClusterSpecMonitor struct {
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceID" tf:"log_analytics_workspace_id"`
	PrimaryKey              *string `json:"-" sensitive:"true" tf:"primary_key"`
}

type InteractiveQueryClusterSpecNetwork struct {
	// +optional
	ConnectionDirection *string `json:"connectionDirection,omitempty" tf:"connection_direction"`
	// +optional
	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled"`
}

type InteractiveQueryClusterSpecRolesHeadNode struct {
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	Username *string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
	VmSize           *string `json:"vmSize" tf:"vm_size"`
}

type InteractiveQueryClusterSpecRolesWorkerNodeAutoscaleCapacity struct {
	MaxInstanceCount *int64 `json:"maxInstanceCount" tf:"max_instance_count"`
	MinInstanceCount *int64 `json:"minInstanceCount" tf:"min_instance_count"`
}

type InteractiveQueryClusterSpecRolesWorkerNodeAutoscaleRecurrenceSchedule struct {
	Days                []string `json:"days" tf:"days"`
	TargetInstanceCount *int64   `json:"targetInstanceCount" tf:"target_instance_count"`
	Time                *string  `json:"time" tf:"time"`
}

type InteractiveQueryClusterSpecRolesWorkerNodeAutoscaleRecurrence struct {
	// +kubebuilder:validation:MinItems=1
	Schedule []InteractiveQueryClusterSpecRolesWorkerNodeAutoscaleRecurrenceSchedule `json:"schedule" tf:"schedule"`
	Timezone *string                                                                 `json:"timezone" tf:"timezone"`
}

type InteractiveQueryClusterSpecRolesWorkerNodeAutoscale struct {
	// +optional
	Capacity *InteractiveQueryClusterSpecRolesWorkerNodeAutoscaleCapacity `json:"capacity,omitempty" tf:"capacity"`
	// +optional
	Recurrence *InteractiveQueryClusterSpecRolesWorkerNodeAutoscaleRecurrence `json:"recurrence,omitempty" tf:"recurrence"`
}

type InteractiveQueryClusterSpecRolesWorkerNode struct {
	// +optional
	Autoscale *InteractiveQueryClusterSpecRolesWorkerNodeAutoscale `json:"autoscale,omitempty" tf:"autoscale"`
	// +optional
	// Deprecated
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`
	// +optional
	SubnetID            *string `json:"subnetID,omitempty" tf:"subnet_id"`
	TargetInstanceCount *int64  `json:"targetInstanceCount" tf:"target_instance_count"`
	Username            *string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
	VmSize           *string `json:"vmSize" tf:"vm_size"`
}

type InteractiveQueryClusterSpecRolesZookeeperNode struct {
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	Username *string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
	VmSize           *string `json:"vmSize" tf:"vm_size"`
}

type InteractiveQueryClusterSpecRoles struct {
	HeadNode      *InteractiveQueryClusterSpecRolesHeadNode      `json:"headNode" tf:"head_node"`
	WorkerNode    *InteractiveQueryClusterSpecRolesWorkerNode    `json:"workerNode" tf:"worker_node"`
	ZookeeperNode *InteractiveQueryClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type InteractiveQueryClusterSpecStorageAccount struct {
	IsDefault          *bool   `json:"isDefault" tf:"is_default"`
	StorageAccountKey  *string `json:"-" sensitive:"true" tf:"storage_account_key"`
	StorageContainerID *string `json:"storageContainerID" tf:"storage_container_id"`
}

type InteractiveQueryClusterSpecStorageAccountGen2 struct {
	FilesystemID              *string `json:"filesystemID" tf:"filesystem_id"`
	IsDefault                 *bool   `json:"isDefault" tf:"is_default"`
	ManagedIdentityResourceID *string `json:"managedIdentityResourceID" tf:"managed_identity_resource_id"`
	StorageResourceID         *string `json:"storageResourceID" tf:"storage_resource_id"`
}

type InteractiveQueryClusterSpec struct {
	State *InteractiveQueryClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource InteractiveQueryClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InteractiveQueryClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterVersion   *string                                      `json:"clusterVersion" tf:"cluster_version"`
	ComponentVersion *InteractiveQueryClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	Gateway          *InteractiveQueryClusterSpecGateway          `json:"gateway" tf:"gateway"`
	// +optional
	HttpsEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint"`
	Location      *string `json:"location" tf:"location"`
	// +optional
	Metastores *InteractiveQueryClusterSpecMetastores `json:"metastores,omitempty" tf:"metastores"`
	// +optional
	Monitor *InteractiveQueryClusterSpecMonitor `json:"monitor,omitempty" tf:"monitor"`
	Name    *string                             `json:"name" tf:"name"`
	// +optional
	Network           *InteractiveQueryClusterSpecNetwork `json:"network,omitempty" tf:"network"`
	ResourceGroupName *string                             `json:"resourceGroupName" tf:"resource_group_name"`
	Roles             *InteractiveQueryClusterSpecRoles   `json:"roles" tf:"roles"`
	// +optional
	SshEndpoint *string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint"`
	// +optional
	StorageAccount []InteractiveQueryClusterSpecStorageAccount `json:"storageAccount,omitempty" tf:"storage_account"`
	// +optional
	StorageAccountGen2 *InteractiveQueryClusterSpecStorageAccountGen2 `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	Tier *string            `json:"tier" tf:"tier"`
	// +optional
	TlsMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`
}

type InteractiveQueryClusterStatus struct {
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

// InteractiveQueryClusterList is a list of InteractiveQueryClusters
type InteractiveQueryClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InteractiveQueryCluster CRD objects
	Items []InteractiveQueryCluster `json:"items,omitempty"`
}

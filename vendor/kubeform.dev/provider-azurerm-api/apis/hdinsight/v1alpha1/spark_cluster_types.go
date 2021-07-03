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

type SparkCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SparkClusterSpec   `json:"spec,omitempty"`
	Status            SparkClusterStatus `json:"status,omitempty"`
}

type SparkClusterSpecComponentVersion struct {
	Spark *string `json:"spark" tf:"spark"`
}

type SparkClusterSpecGateway struct {
	// +optional
	// Deprecated
	Enabled  *bool   `json:"enabled,omitempty" tf:"enabled"`
	Password *string `json:"-" sensitive:"true" tf:"password"`
	Username *string `json:"username" tf:"username"`
}

type SparkClusterSpecMetastoresAmbari struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type SparkClusterSpecMetastoresHive struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type SparkClusterSpecMetastoresOozie struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type SparkClusterSpecMetastores struct {
	// +optional
	Ambari *SparkClusterSpecMetastoresAmbari `json:"ambari,omitempty" tf:"ambari"`
	// +optional
	Hive *SparkClusterSpecMetastoresHive `json:"hive,omitempty" tf:"hive"`
	// +optional
	Oozie *SparkClusterSpecMetastoresOozie `json:"oozie,omitempty" tf:"oozie"`
}

type SparkClusterSpecMonitor struct {
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceID" tf:"log_analytics_workspace_id"`
	PrimaryKey              *string `json:"-" sensitive:"true" tf:"primary_key"`
}

type SparkClusterSpecNetwork struct {
	// +optional
	ConnectionDirection *string `json:"connectionDirection,omitempty" tf:"connection_direction"`
	// +optional
	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled"`
}

type SparkClusterSpecRolesHeadNode struct {
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

type SparkClusterSpecRolesWorkerNodeAutoscaleCapacity struct {
	MaxInstanceCount *int64 `json:"maxInstanceCount" tf:"max_instance_count"`
	MinInstanceCount *int64 `json:"minInstanceCount" tf:"min_instance_count"`
}

type SparkClusterSpecRolesWorkerNodeAutoscaleRecurrenceSchedule struct {
	Days                []string `json:"days" tf:"days"`
	TargetInstanceCount *int64   `json:"targetInstanceCount" tf:"target_instance_count"`
	Time                *string  `json:"time" tf:"time"`
}

type SparkClusterSpecRolesWorkerNodeAutoscaleRecurrence struct {
	// +kubebuilder:validation:MinItems=1
	Schedule []SparkClusterSpecRolesWorkerNodeAutoscaleRecurrenceSchedule `json:"schedule" tf:"schedule"`
	Timezone *string                                                      `json:"timezone" tf:"timezone"`
}

type SparkClusterSpecRolesWorkerNodeAutoscale struct {
	// +optional
	Capacity *SparkClusterSpecRolesWorkerNodeAutoscaleCapacity `json:"capacity,omitempty" tf:"capacity"`
	// +optional
	Recurrence *SparkClusterSpecRolesWorkerNodeAutoscaleRecurrence `json:"recurrence,omitempty" tf:"recurrence"`
}

type SparkClusterSpecRolesWorkerNode struct {
	// +optional
	Autoscale *SparkClusterSpecRolesWorkerNodeAutoscale `json:"autoscale,omitempty" tf:"autoscale"`
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

type SparkClusterSpecRolesZookeeperNode struct {
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

type SparkClusterSpecRoles struct {
	HeadNode      *SparkClusterSpecRolesHeadNode      `json:"headNode" tf:"head_node"`
	WorkerNode    *SparkClusterSpecRolesWorkerNode    `json:"workerNode" tf:"worker_node"`
	ZookeeperNode *SparkClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type SparkClusterSpecStorageAccount struct {
	IsDefault          *bool   `json:"isDefault" tf:"is_default"`
	StorageAccountKey  *string `json:"-" sensitive:"true" tf:"storage_account_key"`
	StorageContainerID *string `json:"storageContainerID" tf:"storage_container_id"`
}

type SparkClusterSpecStorageAccountGen2 struct {
	FilesystemID              *string `json:"filesystemID" tf:"filesystem_id"`
	IsDefault                 *bool   `json:"isDefault" tf:"is_default"`
	ManagedIdentityResourceID *string `json:"managedIdentityResourceID" tf:"managed_identity_resource_id"`
	StorageResourceID         *string `json:"storageResourceID" tf:"storage_resource_id"`
}

type SparkClusterSpec struct {
	KubeformOutput *SparkClusterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource SparkClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type SparkClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterVersion   *string                           `json:"clusterVersion" tf:"cluster_version"`
	ComponentVersion *SparkClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	Gateway          *SparkClusterSpecGateway          `json:"gateway" tf:"gateway"`
	// +optional
	HttpsEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint"`
	Location      *string `json:"location" tf:"location"`
	// +optional
	Metastores *SparkClusterSpecMetastores `json:"metastores,omitempty" tf:"metastores"`
	// +optional
	Monitor *SparkClusterSpecMonitor `json:"monitor,omitempty" tf:"monitor"`
	Name    *string                  `json:"name" tf:"name"`
	// +optional
	Network           *SparkClusterSpecNetwork `json:"network,omitempty" tf:"network"`
	ResourceGroupName *string                  `json:"resourceGroupName" tf:"resource_group_name"`
	Roles             *SparkClusterSpecRoles   `json:"roles" tf:"roles"`
	// +optional
	SshEndpoint *string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint"`
	// +optional
	StorageAccount []SparkClusterSpecStorageAccount `json:"storageAccount,omitempty" tf:"storage_account"`
	// +optional
	StorageAccountGen2 *SparkClusterSpecStorageAccountGen2 `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	Tier *string            `json:"tier" tf:"tier"`
	// +optional
	TlsMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`
}

type SparkClusterStatus struct {
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

// SparkClusterList is a list of SparkClusters
type SparkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SparkCluster CRD objects
	Items []SparkCluster `json:"items,omitempty"`
}
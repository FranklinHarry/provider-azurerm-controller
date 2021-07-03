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

type CloudJavaDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudJavaDeploymentSpec   `json:"spec,omitempty"`
	Status            CloudJavaDeploymentStatus `json:"status,omitempty"`
}

type CloudJavaDeploymentSpec struct {
	KubeformOutput *CloudJavaDeploymentSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource CloudJavaDeploymentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CloudJavaDeploymentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Cpu *int64 `json:"cpu,omitempty" tf:"cpu"`
	// +optional
	EnvironmentVariables *map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables"`
	// +optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count"`
	// +optional
	JvmOptions *string `json:"jvmOptions,omitempty" tf:"jvm_options"`
	// +optional
	MemoryInGb *int64  `json:"memoryInGb,omitempty" tf:"memory_in_gb"`
	Name       *string `json:"name" tf:"name"`
	// +optional
	RuntimeVersion   *string `json:"runtimeVersion,omitempty" tf:"runtime_version"`
	SpringCloudAppID *string `json:"springCloudAppID" tf:"spring_cloud_app_id"`
}

type CloudJavaDeploymentStatus struct {
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

// CloudJavaDeploymentList is a list of CloudJavaDeployments
type CloudJavaDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudJavaDeployment CRD objects
	Items []CloudJavaDeployment `json:"items,omitempty"`
}

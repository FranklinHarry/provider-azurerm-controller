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

type CustomDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDomainSpec   `json:"spec,omitempty"`
	Status            CustomDomainStatus `json:"status,omitempty"`
}

type CustomDomainSpecDeveloperPortal struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	HostName            *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
}

type CustomDomainSpecManagement struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	HostName            *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
}

type CustomDomainSpecPortal struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	HostName            *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
}

type CustomDomainSpecProxy struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	// +optional
	DefaultSslBinding *bool   `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding"`
	HostName          *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
}

type CustomDomainSpecScm struct {
	// +optional
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificatePassword *string `json:"-" sensitive:"true" tf:"certificate_password"`
	HostName            *string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID *string `json:"keyVaultID,omitempty" tf:"key_vault_id"`
	// +optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate"`
}

type CustomDomainSpec struct {
	KubeformOutput *CustomDomainSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource CustomDomainSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type CustomDomainSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementID *string `json:"apiManagementID" tf:"api_management_id"`
	// +optional
	DeveloperPortal []CustomDomainSpecDeveloperPortal `json:"developerPortal,omitempty" tf:"developer_portal"`
	// +optional
	Management []CustomDomainSpecManagement `json:"management,omitempty" tf:"management"`
	// +optional
	Portal []CustomDomainSpecPortal `json:"portal,omitempty" tf:"portal"`
	// +optional
	Proxy []CustomDomainSpecProxy `json:"proxy,omitempty" tf:"proxy"`
	// +optional
	Scm []CustomDomainSpecScm `json:"scm,omitempty" tf:"scm"`
}

type CustomDomainStatus struct {
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

// CustomDomainList is a list of CustomDomains
type CustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomDomain CRD objects
	Items []CustomDomain `json:"items,omitempty"`
}

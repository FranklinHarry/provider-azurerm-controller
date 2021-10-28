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

type Api struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiSpec   `json:"spec,omitempty"`
	Status            ApiStatus `json:"status,omitempty"`
}

type ApiSpecImportWsdlSelector struct {
	EndpointName *string `json:"endpointName" tf:"endpoint_name"`
	ServiceName  *string `json:"serviceName" tf:"service_name"`
}

type ApiSpecImport struct {
	ContentFormat *string `json:"contentFormat" tf:"content_format"`
	ContentValue  *string `json:"contentValue" tf:"content_value"`
	// +optional
	WsdlSelector *ApiSpecImportWsdlSelector `json:"wsdlSelector,omitempty" tf:"wsdl_selector"`
}

type ApiSpecOauth2Authorization struct {
	AuthorizationServerName *string `json:"authorizationServerName" tf:"authorization_server_name"`
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
}

type ApiSpecOpenidAuthentication struct {
	// +optional
	BearerTokenSendingMethods []string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods"`
	OpenidProviderName        *string  `json:"openidProviderName" tf:"openid_provider_name"`
}

type ApiSpecSubscriptionKeyParameterNames struct {
	Header *string `json:"header" tf:"header"`
	Query  *string `json:"query" tf:"query"`
}

type ApiSpec struct {
	State *ApiSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApiSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApiSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementName *string `json:"apiManagementName" tf:"api_management_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	Import *ApiSpecImport `json:"import,omitempty" tf:"import"`
	// +optional
	IsCurrent *bool `json:"isCurrent,omitempty" tf:"is_current"`
	// +optional
	IsOnline *bool   `json:"isOnline,omitempty" tf:"is_online"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	Oauth2Authorization *ApiSpecOauth2Authorization `json:"oauth2Authorization,omitempty" tf:"oauth2_authorization"`
	// +optional
	OpenidAuthentication *ApiSpecOpenidAuthentication `json:"openidAuthentication,omitempty" tf:"openid_authentication"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Protocols         []string `json:"protocols,omitempty" tf:"protocols"`
	ResourceGroupName *string  `json:"resourceGroupName" tf:"resource_group_name"`
	Revision          *string  `json:"revision" tf:"revision"`
	// +optional
	RevisionDescription *string `json:"revisionDescription,omitempty" tf:"revision_description"`
	// +optional
	ServiceURL *string `json:"serviceURL,omitempty" tf:"service_url"`
	// +optional
	SoapPassThrough *bool `json:"soapPassThrough,omitempty" tf:"soap_pass_through"`
	// +optional
	SourceAPIID *string `json:"sourceAPIID,omitempty" tf:"source_api_id"`
	// +optional
	SubscriptionKeyParameterNames *ApiSpecSubscriptionKeyParameterNames `json:"subscriptionKeyParameterNames,omitempty" tf:"subscription_key_parameter_names"`
	// +optional
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty" tf:"subscription_required"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VersionDescription *string `json:"versionDescription,omitempty" tf:"version_description"`
	// +optional
	VersionSetID *string `json:"versionSetID,omitempty" tf:"version_set_id"`
}

type ApiStatus struct {
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

// ApiList is a list of Apis
type ApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Api CRD objects
	Items []Api `json:"items,omitempty"`
}

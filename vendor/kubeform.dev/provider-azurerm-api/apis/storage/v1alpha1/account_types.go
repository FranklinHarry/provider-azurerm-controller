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

type Account struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec,omitempty"`
	Status            AccountStatus `json:"status,omitempty"`
}

type AccountSpecAzureFilesAuthenticationActiveDirectory struct {
	DomainGuid        *string `json:"domainGuid" tf:"domain_guid"`
	DomainName        *string `json:"domainName" tf:"domain_name"`
	DomainSid         *string `json:"domainSid" tf:"domain_sid"`
	ForestName        *string `json:"forestName" tf:"forest_name"`
	NetbiosDomainName *string `json:"netbiosDomainName" tf:"netbios_domain_name"`
	StorageSid        *string `json:"storageSid" tf:"storage_sid"`
}

type AccountSpecAzureFilesAuthentication struct {
	// +optional
	ActiveDirectory *AccountSpecAzureFilesAuthenticationActiveDirectory `json:"activeDirectory,omitempty" tf:"active_directory"`
	DirectoryType   *string                                             `json:"directoryType" tf:"directory_type"`
}

type AccountSpecBlobPropertiesContainerDeleteRetentionPolicy struct {
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
}

type AccountSpecBlobPropertiesCorsRule struct {
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	AllowedHeaders []string `json:"allowedHeaders" tf:"allowed_headers"`
	// +kubebuilder:validation:MaxItems=64
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// +kubebuilder:validation:MaxItems=64
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	ExposedHeaders  []string `json:"exposedHeaders" tf:"exposed_headers"`
	MaxAgeInSeconds *int64   `json:"maxAgeInSeconds" tf:"max_age_in_seconds"`
}

type AccountSpecBlobPropertiesDeleteRetentionPolicy struct {
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
}

type AccountSpecBlobProperties struct {
	// +optional
	ChangeFeedEnabled *bool `json:"changeFeedEnabled,omitempty" tf:"change_feed_enabled"`
	// +optional
	ContainerDeleteRetentionPolicy *AccountSpecBlobPropertiesContainerDeleteRetentionPolicy `json:"containerDeleteRetentionPolicy,omitempty" tf:"container_delete_retention_policy"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	CorsRule []AccountSpecBlobPropertiesCorsRule `json:"corsRule,omitempty" tf:"cors_rule"`
	// +optional
	DefaultServiceVersion *string `json:"defaultServiceVersion,omitempty" tf:"default_service_version"`
	// +optional
	DeleteRetentionPolicy *AccountSpecBlobPropertiesDeleteRetentionPolicy `json:"deleteRetentionPolicy,omitempty" tf:"delete_retention_policy"`
	// +optional
	LastAccessTimeEnabled *bool `json:"lastAccessTimeEnabled,omitempty" tf:"last_access_time_enabled"`
	// +optional
	VersioningEnabled *bool `json:"versioningEnabled,omitempty" tf:"versioning_enabled"`
}

type AccountSpecCustomDomain struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	UseSubdomain *bool `json:"useSubdomain,omitempty" tf:"use_subdomain"`
}

type AccountSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type AccountSpecNetworkRulesPrivateLinkAccess struct {
	EndpointResourceID *string `json:"endpointResourceID" tf:"endpoint_resource_id"`
	// +optional
	EndpointTenantID *string `json:"endpointTenantID,omitempty" tf:"endpoint_tenant_id"`
}

type AccountSpecNetworkRules struct {
	// +optional
	Bypass        []string `json:"bypass,omitempty" tf:"bypass"`
	DefaultAction *string  `json:"defaultAction" tf:"default_action"`
	// +optional
	IpRules []string `json:"ipRules,omitempty" tf:"ip_rules"`
	// +optional
	PrivateLinkAccess []AccountSpecNetworkRulesPrivateLinkAccess `json:"privateLinkAccess,omitempty" tf:"private_link_access"`
	// +optional
	VirtualNetworkSubnetIDS []string `json:"virtualNetworkSubnetIDS,omitempty" tf:"virtual_network_subnet_ids"`
}

type AccountSpecQueuePropertiesCorsRule struct {
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	AllowedHeaders []string `json:"allowedHeaders" tf:"allowed_headers"`
	// +kubebuilder:validation:MaxItems=64
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// +kubebuilder:validation:MaxItems=64
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	ExposedHeaders  []string `json:"exposedHeaders" tf:"exposed_headers"`
	MaxAgeInSeconds *int64   `json:"maxAgeInSeconds" tf:"max_age_in_seconds"`
}

type AccountSpecQueuePropertiesHourMetrics struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	IncludeApis *bool `json:"includeApis,omitempty" tf:"include_apis"`
	// +optional
	RetentionPolicyDays *int64  `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days"`
	Version             *string `json:"version" tf:"version"`
}

type AccountSpecQueuePropertiesLogging struct {
	Delete *bool `json:"delete" tf:"delete"`
	Read   *bool `json:"read" tf:"read"`
	// +optional
	RetentionPolicyDays *int64  `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days"`
	Version             *string `json:"version" tf:"version"`
	Write               *bool   `json:"write" tf:"write"`
}

type AccountSpecQueuePropertiesMinuteMetrics struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	IncludeApis *bool `json:"includeApis,omitempty" tf:"include_apis"`
	// +optional
	RetentionPolicyDays *int64  `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days"`
	Version             *string `json:"version" tf:"version"`
}

type AccountSpecQueueProperties struct {
	// +optional
	// +kubebuilder:validation:MaxItems=5
	CorsRule []AccountSpecQueuePropertiesCorsRule `json:"corsRule,omitempty" tf:"cors_rule"`
	// +optional
	HourMetrics *AccountSpecQueuePropertiesHourMetrics `json:"hourMetrics,omitempty" tf:"hour_metrics"`
	// +optional
	Logging *AccountSpecQueuePropertiesLogging `json:"logging,omitempty" tf:"logging"`
	// +optional
	MinuteMetrics *AccountSpecQueuePropertiesMinuteMetrics `json:"minuteMetrics,omitempty" tf:"minute_metrics"`
}

type AccountSpecRouting struct {
	// +optional
	Choice *string `json:"choice,omitempty" tf:"choice"`
	// +optional
	PublishInternetEndpoints *bool `json:"publishInternetEndpoints,omitempty" tf:"publish_internet_endpoints"`
	// +optional
	PublishMicrosoftEndpoints *bool `json:"publishMicrosoftEndpoints,omitempty" tf:"publish_microsoft_endpoints"`
}

type AccountSpecSharePropertiesCorsRule struct {
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	AllowedHeaders []string `json:"allowedHeaders" tf:"allowed_headers"`
	// +kubebuilder:validation:MaxItems=64
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// +kubebuilder:validation:MaxItems=64
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +kubebuilder:validation:MaxItems=64
	// +kubebuilder:validation:MinItems=1
	ExposedHeaders  []string `json:"exposedHeaders" tf:"exposed_headers"`
	MaxAgeInSeconds *int64   `json:"maxAgeInSeconds" tf:"max_age_in_seconds"`
}

type AccountSpecSharePropertiesRetentionPolicy struct {
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
}

type AccountSpecSharePropertiesSmb struct {
	// +optional
	AuthenticationTypes []string `json:"authenticationTypes,omitempty" tf:"authentication_types"`
	// +optional
	ChannelEncryptionType []string `json:"channelEncryptionType,omitempty" tf:"channel_encryption_type"`
	// +optional
	KerberosTicketEncryptionType []string `json:"kerberosTicketEncryptionType,omitempty" tf:"kerberos_ticket_encryption_type"`
	// +optional
	Versions []string `json:"versions,omitempty" tf:"versions"`
}

type AccountSpecShareProperties struct {
	// +optional
	// +kubebuilder:validation:MaxItems=5
	CorsRule []AccountSpecSharePropertiesCorsRule `json:"corsRule,omitempty" tf:"cors_rule"`
	// +optional
	RetentionPolicy *AccountSpecSharePropertiesRetentionPolicy `json:"retentionPolicy,omitempty" tf:"retention_policy"`
	// +optional
	Smb *AccountSpecSharePropertiesSmb `json:"smb,omitempty" tf:"smb"`
}

type AccountSpecStaticWebsite struct {
	// +optional
	Error404Document *string `json:"error404Document,omitempty" tf:"error_404_document"`
	// +optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document"`
}

type AccountSpec struct {
	State *AccountSpecResource `json:"state,omitempty" tf:"-"`

	Resource AccountSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AccountSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier"`
	// +optional
	AccountKind            *string `json:"accountKind,omitempty" tf:"account_kind"`
	AccountReplicationType *string `json:"accountReplicationType" tf:"account_replication_type"`
	AccountTier            *string `json:"accountTier" tf:"account_tier"`
	// +optional
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty" tf:"allow_blob_public_access"`
	// +optional
	AzureFilesAuthentication *AccountSpecAzureFilesAuthentication `json:"azureFilesAuthentication,omitempty" tf:"azure_files_authentication"`
	// +optional
	BlobProperties *AccountSpecBlobProperties `json:"blobProperties,omitempty" tf:"blob_properties"`
	// +optional
	CustomDomain *AccountSpecCustomDomain `json:"customDomain,omitempty" tf:"custom_domain"`
	// +optional
	EnableHTTPSTrafficOnly *bool `json:"enableHTTPSTrafficOnly,omitempty" tf:"enable_https_traffic_only"`
	// +optional
	Identity *AccountSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty" tf:"is_hns_enabled"`
	// +optional
	LargeFileShareEnabled *bool   `json:"largeFileShareEnabled,omitempty" tf:"large_file_share_enabled"`
	Location              *string `json:"location" tf:"location"`
	// +optional
	MinTlsVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version"`
	Name          *string `json:"name" tf:"name"`
	// +optional
	NetworkRules *AccountSpecNetworkRules `json:"networkRules,omitempty" tf:"network_rules"`
	// +optional
	Nfsv3Enabled *bool `json:"nfsv3Enabled,omitempty" tf:"nfsv3_enabled"`
	// +optional
	PrimaryAccessKey *string `json:"-" sensitive:"true" tf:"primary_access_key"`
	// +optional
	PrimaryBlobConnectionString *string `json:"-" sensitive:"true" tf:"primary_blob_connection_string"`
	// +optional
	PrimaryBlobEndpoint *string `json:"primaryBlobEndpoint,omitempty" tf:"primary_blob_endpoint"`
	// +optional
	PrimaryBlobHost *string `json:"primaryBlobHost,omitempty" tf:"primary_blob_host"`
	// +optional
	PrimaryConnectionString *string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	// +optional
	PrimaryDfsEndpoint *string `json:"primaryDfsEndpoint,omitempty" tf:"primary_dfs_endpoint"`
	// +optional
	PrimaryDfsHost *string `json:"primaryDfsHost,omitempty" tf:"primary_dfs_host"`
	// +optional
	PrimaryFileEndpoint *string `json:"primaryFileEndpoint,omitempty" tf:"primary_file_endpoint"`
	// +optional
	PrimaryFileHost *string `json:"primaryFileHost,omitempty" tf:"primary_file_host"`
	// +optional
	PrimaryLocation *string `json:"primaryLocation,omitempty" tf:"primary_location"`
	// +optional
	PrimaryQueueEndpoint *string `json:"primaryQueueEndpoint,omitempty" tf:"primary_queue_endpoint"`
	// +optional
	PrimaryQueueHost *string `json:"primaryQueueHost,omitempty" tf:"primary_queue_host"`
	// +optional
	PrimaryTableEndpoint *string `json:"primaryTableEndpoint,omitempty" tf:"primary_table_endpoint"`
	// +optional
	PrimaryTableHost *string `json:"primaryTableHost,omitempty" tf:"primary_table_host"`
	// +optional
	PrimaryWebEndpoint *string `json:"primaryWebEndpoint,omitempty" tf:"primary_web_endpoint"`
	// +optional
	PrimaryWebHost *string `json:"primaryWebHost,omitempty" tf:"primary_web_host"`
	// +optional
	QueueProperties   *AccountSpecQueueProperties `json:"queueProperties,omitempty" tf:"queue_properties"`
	ResourceGroupName *string                     `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Routing *AccountSpecRouting `json:"routing,omitempty" tf:"routing"`
	// +optional
	SecondaryAccessKey *string `json:"-" sensitive:"true" tf:"secondary_access_key"`
	// +optional
	SecondaryBlobConnectionString *string `json:"-" sensitive:"true" tf:"secondary_blob_connection_string"`
	// +optional
	SecondaryBlobEndpoint *string `json:"secondaryBlobEndpoint,omitempty" tf:"secondary_blob_endpoint"`
	// +optional
	SecondaryBlobHost *string `json:"secondaryBlobHost,omitempty" tf:"secondary_blob_host"`
	// +optional
	SecondaryConnectionString *string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	// +optional
	SecondaryDfsEndpoint *string `json:"secondaryDfsEndpoint,omitempty" tf:"secondary_dfs_endpoint"`
	// +optional
	SecondaryDfsHost *string `json:"secondaryDfsHost,omitempty" tf:"secondary_dfs_host"`
	// +optional
	SecondaryFileEndpoint *string `json:"secondaryFileEndpoint,omitempty" tf:"secondary_file_endpoint"`
	// +optional
	SecondaryFileHost *string `json:"secondaryFileHost,omitempty" tf:"secondary_file_host"`
	// +optional
	SecondaryLocation *string `json:"secondaryLocation,omitempty" tf:"secondary_location"`
	// +optional
	SecondaryQueueEndpoint *string `json:"secondaryQueueEndpoint,omitempty" tf:"secondary_queue_endpoint"`
	// +optional
	SecondaryQueueHost *string `json:"secondaryQueueHost,omitempty" tf:"secondary_queue_host"`
	// +optional
	SecondaryTableEndpoint *string `json:"secondaryTableEndpoint,omitempty" tf:"secondary_table_endpoint"`
	// +optional
	SecondaryTableHost *string `json:"secondaryTableHost,omitempty" tf:"secondary_table_host"`
	// +optional
	SecondaryWebEndpoint *string `json:"secondaryWebEndpoint,omitempty" tf:"secondary_web_endpoint"`
	// +optional
	SecondaryWebHost *string `json:"secondaryWebHost,omitempty" tf:"secondary_web_host"`
	// +optional
	ShareProperties *AccountSpecShareProperties `json:"shareProperties,omitempty" tf:"share_properties"`
	// +optional
	StaticWebsite *AccountSpecStaticWebsite `json:"staticWebsite,omitempty" tf:"static_website"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type AccountStatus struct {
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

// AccountList is a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Account CRD objects
	Items []Account `json:"items,omitempty"`
}

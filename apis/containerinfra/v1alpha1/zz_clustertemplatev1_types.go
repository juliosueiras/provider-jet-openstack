/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClustertemplateV1Observation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type ClustertemplateV1Parameters struct {

	// +kubebuilder:validation:Optional
	ApiserverPort *float64 `json:"apiserverPort,omitempty" tf:"apiserver_port,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterDistro *string `json:"clusterDistro,omitempty" tf:"cluster_distro,omitempty"`

	// +kubebuilder:validation:Required
	Coe *string `json:"coe" tf:"coe,omitempty"`

	// +kubebuilder:validation:Optional
	DNSNameserver *string `json:"dnsNameserver,omitempty" tf:"dns_nameserver,omitempty"`

	// +kubebuilder:validation:Optional
	DockerStorageDriver *string `json:"dockerStorageDriver,omitempty" tf:"docker_storage_driver,omitempty"`

	// +kubebuilder:validation:Optional
	DockerVolumeSize *float64 `json:"dockerVolumeSize,omitempty" tf:"docker_volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// +kubebuilder:validation:Optional
	FixedNetwork *string `json:"fixedNetwork,omitempty" tf:"fixed_network,omitempty"`

	// +kubebuilder:validation:Optional
	FixedSubnet *string `json:"fixedSubnet,omitempty" tf:"fixed_subnet,omitempty"`

	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// +kubebuilder:validation:Optional
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPProxy *string `json:"httpProxy,omitempty" tf:"http_proxy,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPSProxy *string `json:"httpsProxy,omitempty" tf:"https_proxy,omitempty"`

	// +kubebuilder:validation:Required
	Image *string `json:"image" tf:"image,omitempty"`

	// +kubebuilder:validation:Optional
	InsecureRegistry *string `json:"insecureRegistry,omitempty" tf:"insecure_registry,omitempty"`

	// +kubebuilder:validation:Optional
	KeypairID *string `json:"keypairId,omitempty" tf:"keypair_id,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// +kubebuilder:validation:Optional
	MasterLBEnabled *bool `json:"masterLbEnabled,omitempty" tf:"master_lb_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkDriver *string `json:"networkDriver,omitempty" tf:"network_driver,omitempty"`

	// +kubebuilder:validation:Optional
	NoProxy *string `json:"noProxy,omitempty" tf:"no_proxy,omitempty"`

	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryEnabled *bool `json:"registryEnabled,omitempty" tf:"registry_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ServerType *string `json:"serverType,omitempty" tf:"server_type,omitempty"`

	// +kubebuilder:validation:Optional
	TLSDisabled *bool `json:"tlsDisabled,omitempty" tf:"tls_disabled,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeDriver *string `json:"volumeDriver,omitempty" tf:"volume_driver,omitempty"`
}

// ClustertemplateV1Spec defines the desired state of ClustertemplateV1
type ClustertemplateV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClustertemplateV1Parameters `json:"forProvider"`
}

// ClustertemplateV1Status defines the observed state of ClustertemplateV1.
type ClustertemplateV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClustertemplateV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClustertemplateV1 is the Schema for the ClustertemplateV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type ClustertemplateV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClustertemplateV1Spec   `json:"spec"`
	Status            ClustertemplateV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClustertemplateV1List contains a list of ClustertemplateV1s
type ClustertemplateV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClustertemplateV1 `json:"items"`
}

// Repository type metadata.
var (
	ClustertemplateV1_Kind             = "ClustertemplateV1"
	ClustertemplateV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClustertemplateV1_Kind}.String()
	ClustertemplateV1_KindAPIVersion   = ClustertemplateV1_Kind + "." + CRDGroupVersion.String()
	ClustertemplateV1_GroupVersionKind = CRDGroupVersion.WithKind(ClustertemplateV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ClustertemplateV1{}, &ClustertemplateV1List{})
}

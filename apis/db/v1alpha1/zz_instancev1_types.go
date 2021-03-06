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

type DatabaseObservation struct {
}

type DatabaseParameters struct {

	// +kubebuilder:validation:Optional
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// +kubebuilder:validation:Optional
	Collate *string `json:"collate,omitempty" tf:"collate,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type InstanceV1DatastoreObservation struct {
}

type InstanceV1DatastoreParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type InstanceV1Observation struct {
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InstanceV1Parameters struct {

	// +kubebuilder:validation:Optional
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// +kubebuilder:validation:Required
	Datastore []InstanceV1DatastoreParameters `json:"datastore" tf:"datastore,omitempty"`

	// +kubebuilder:validation:Optional
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	User []UserParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type NetworkObservation struct {
}

type NetworkParameters struct {

	// +kubebuilder:validation:Optional
	FixedIPV4 *string `json:"fixedIpV4,omitempty" tf:"fixed_ip_v4,omitempty"`

	// +kubebuilder:validation:Optional
	FixedIPV6 *string `json:"fixedIpV6,omitempty" tf:"fixed_ip_v6,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type UserObservation struct {
}

type UserParameters struct {

	// +kubebuilder:validation:Optional
	Databases []*string `json:"databases,omitempty" tf:"databases,omitempty"`

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`
}

// InstanceV1Spec defines the desired state of InstanceV1
type InstanceV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceV1Parameters `json:"forProvider"`
}

// InstanceV1Status defines the observed state of InstanceV1.
type InstanceV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV1 is the Schema for the InstanceV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type InstanceV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceV1Spec   `json:"spec"`
	Status            InstanceV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceV1List contains a list of InstanceV1s
type InstanceV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceV1 `json:"items"`
}

// Repository type metadata.
var (
	InstanceV1_Kind             = "InstanceV1"
	InstanceV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceV1_Kind}.String()
	InstanceV1_KindAPIVersion   = InstanceV1_Kind + "." + CRDGroupVersion.String()
	InstanceV1_GroupVersionKind = CRDGroupVersion.WithKind(InstanceV1_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceV1{}, &InstanceV1List{})
}

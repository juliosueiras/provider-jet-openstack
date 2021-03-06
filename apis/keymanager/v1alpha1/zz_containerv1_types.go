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

type ACLObservation struct {
}

type ACLParameters struct {

	// +kubebuilder:validation:Optional
	Read []ReadParameters `json:"read,omitempty" tf:"read,omitempty"`
}

type ConsumersObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ConsumersParameters struct {
}

type ContainerV1Observation struct {
	Consumers []ConsumersObservation `json:"consumers,omitempty" tf:"consumers,omitempty"`

	ContainerRef *string `json:"containerRef,omitempty" tf:"container_ref,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	CreatorID *string `json:"creatorId,omitempty" tf:"creator_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ContainerV1Parameters struct {

	// +kubebuilder:validation:Optional
	ACL []ACLParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	SecretRefs []SecretRefsParameters `json:"secretRefs,omitempty" tf:"secret_refs,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ReadObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ReadParameters struct {

	// +kubebuilder:validation:Optional
	ProjectAccess *bool `json:"projectAccess,omitempty" tf:"project_access,omitempty"`

	// +kubebuilder:validation:Optional
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type SecretRefsObservation struct {
}

type SecretRefsParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SecretRef *string `json:"secretRef" tf:"secret_ref,omitempty"`
}

// ContainerV1Spec defines the desired state of ContainerV1
type ContainerV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerV1Parameters `json:"forProvider"`
}

// ContainerV1Status defines the observed state of ContainerV1.
type ContainerV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerV1 is the Schema for the ContainerV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type ContainerV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerV1Spec   `json:"spec"`
	Status            ContainerV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerV1List contains a list of ContainerV1s
type ContainerV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerV1 `json:"items"`
}

// Repository type metadata.
var (
	ContainerV1_Kind             = "ContainerV1"
	ContainerV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerV1_Kind}.String()
	ContainerV1_KindAPIVersion   = ContainerV1_Kind + "." + CRDGroupVersion.String()
	ContainerV1_GroupVersionKind = CRDGroupVersion.WithKind(ContainerV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerV1{}, &ContainerV1List{})
}

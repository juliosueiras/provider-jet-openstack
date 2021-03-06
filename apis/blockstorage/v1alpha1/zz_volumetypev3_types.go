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

type VolumeTypeV3Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VolumeTypeV3Parameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraSpecs map[string]*string `json:"extraSpecs,omitempty" tf:"extra_specs,omitempty"`

	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// VolumeTypeV3Spec defines the desired state of VolumeTypeV3
type VolumeTypeV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeTypeV3Parameters `json:"forProvider"`
}

// VolumeTypeV3Status defines the observed state of VolumeTypeV3.
type VolumeTypeV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeTypeV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeTypeV3 is the Schema for the VolumeTypeV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type VolumeTypeV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeTypeV3Spec   `json:"spec"`
	Status            VolumeTypeV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeTypeV3List contains a list of VolumeTypeV3s
type VolumeTypeV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeTypeV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeTypeV3_Kind             = "VolumeTypeV3"
	VolumeTypeV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeTypeV3_Kind}.String()
	VolumeTypeV3_KindAPIVersion   = VolumeTypeV3_Kind + "." + CRDGroupVersion.String()
	VolumeTypeV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeTypeV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeTypeV3{}, &VolumeTypeV3List{})
}

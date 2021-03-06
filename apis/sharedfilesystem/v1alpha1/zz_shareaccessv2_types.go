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

type ShareAccessV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ShareAccessV2Parameters struct {

	// +kubebuilder:validation:Required
	AccessLevel *string `json:"accessLevel" tf:"access_level,omitempty"`

	// +kubebuilder:validation:Required
	AccessTo *string `json:"accessTo" tf:"access_to,omitempty"`

	// +kubebuilder:validation:Required
	AccessType *string `json:"accessType" tf:"access_type,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	ShareID *string `json:"shareId" tf:"share_id,omitempty"`
}

// ShareAccessV2Spec defines the desired state of ShareAccessV2
type ShareAccessV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareAccessV2Parameters `json:"forProvider"`
}

// ShareAccessV2Status defines the observed state of ShareAccessV2.
type ShareAccessV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareAccessV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ShareAccessV2 is the Schema for the ShareAccessV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type ShareAccessV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShareAccessV2Spec   `json:"spec"`
	Status            ShareAccessV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareAccessV2List contains a list of ShareAccessV2s
type ShareAccessV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShareAccessV2 `json:"items"`
}

// Repository type metadata.
var (
	ShareAccessV2_Kind             = "ShareAccessV2"
	ShareAccessV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ShareAccessV2_Kind}.String()
	ShareAccessV2_KindAPIVersion   = ShareAccessV2_Kind + "." + CRDGroupVersion.String()
	ShareAccessV2_GroupVersionKind = CRDGroupVersion.WithKind(ShareAccessV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ShareAccessV2{}, &ShareAccessV2List{})
}

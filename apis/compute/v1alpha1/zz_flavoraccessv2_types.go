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

type FlavorAccessV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FlavorAccessV2Parameters struct {

	// +kubebuilder:validation:Required
	FlavorID *string `json:"flavorId" tf:"flavor_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// FlavorAccessV2Spec defines the desired state of FlavorAccessV2
type FlavorAccessV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlavorAccessV2Parameters `json:"forProvider"`
}

// FlavorAccessV2Status defines the observed state of FlavorAccessV2.
type FlavorAccessV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlavorAccessV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlavorAccessV2 is the Schema for the FlavorAccessV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type FlavorAccessV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlavorAccessV2Spec   `json:"spec"`
	Status            FlavorAccessV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlavorAccessV2List contains a list of FlavorAccessV2s
type FlavorAccessV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlavorAccessV2 `json:"items"`
}

// Repository type metadata.
var (
	FlavorAccessV2_Kind             = "FlavorAccessV2"
	FlavorAccessV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlavorAccessV2_Kind}.String()
	FlavorAccessV2_KindAPIVersion   = FlavorAccessV2_Kind + "." + CRDGroupVersion.String()
	FlavorAccessV2_GroupVersionKind = CRDGroupVersion.WithKind(FlavorAccessV2_Kind)
)

func init() {
	SchemeBuilder.Register(&FlavorAccessV2{}, &FlavorAccessV2List{})
}

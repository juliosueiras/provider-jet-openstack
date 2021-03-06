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

type L7PolicyV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type L7PolicyV2Parameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	ListenerID *string `json:"listenerId" tf:"listener_id,omitempty"`

	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// +kubebuilder:validation:Optional
	RedirectPoolID *string `json:"redirectPoolId,omitempty" tf:"redirect_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// L7PolicyV2Spec defines the desired state of L7PolicyV2
type L7PolicyV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     L7PolicyV2Parameters `json:"forProvider"`
}

// L7PolicyV2Status defines the observed state of L7PolicyV2.
type L7PolicyV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        L7PolicyV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// L7PolicyV2 is the Schema for the L7PolicyV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type L7PolicyV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              L7PolicyV2Spec   `json:"spec"`
	Status            L7PolicyV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// L7PolicyV2List contains a list of L7PolicyV2s
type L7PolicyV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L7PolicyV2 `json:"items"`
}

// Repository type metadata.
var (
	L7PolicyV2_Kind             = "L7PolicyV2"
	L7PolicyV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: L7PolicyV2_Kind}.String()
	L7PolicyV2_KindAPIVersion   = L7PolicyV2_Kind + "." + CRDGroupVersion.String()
	L7PolicyV2_GroupVersionKind = CRDGroupVersion.WithKind(L7PolicyV2_Kind)
)

func init() {
	SchemeBuilder.Register(&L7PolicyV2{}, &L7PolicyV2List{})
}

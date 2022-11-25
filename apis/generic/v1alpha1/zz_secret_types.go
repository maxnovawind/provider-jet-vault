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

type SecretObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecretParameters struct {

	// JSON-encoded secret data to write.
	// +kubebuilder:validation:Required
	DataJSON *string `json:"dataJson" tf:"data_json,omitempty"`

	// Only applicable for kv-v2 stores. If set, permanently deletes all versions for the specified key.
	// +kubebuilder:validation:Optional
	DeleteAllVersions *bool `json:"deleteAllVersions,omitempty" tf:"delete_all_versions,omitempty"`

	// Don't attempt to read the token from Vault if true; drift won't be detected.
	// +kubebuilder:validation:Optional
	DisableRead *bool `json:"disableRead,omitempty" tf:"disable_read,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Full path where the generic secret will be written.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretParameters `json:"forProvider"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vaultjet}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec"`
	Status            SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	Secret_Kind             = "Secret"
	Secret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Secret_Kind}.String()
	Secret_KindAPIVersion   = Secret_Kind + "." + CRDGroupVersion.String()
	Secret_GroupVersionKind = CRDGroupVersion.WithKind(Secret_Kind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}

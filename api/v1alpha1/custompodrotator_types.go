/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomPodRotatorSpec defines the desired state of CustomPodRotator
type CustomPodRotatorSpec struct {
	TargetDeployment string `json:"targetDeployment"`   // Name of the deployment to rotate
	Schedule         string `json:"schedule"`           // Cron schedule e.g. "0 3 * * *"
	TimeZone         string `json:"timeZone,omitempty"` // Timezone e.g. "UTC"
}

// CustomPodRotatorStatus defines the observed state of CustomPodRotator
type CustomPodRotatorStatus struct {
	LastRestartTime metav1.Time `json:"lastRestartTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CustomPodRotator is the Schema for the custompodrotators API
type CustomPodRotator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomPodRotatorSpec   `json:"spec,omitempty"`
	Status CustomPodRotatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomPodRotatorList contains a list of CustomPodRotator
type CustomPodRotatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomPodRotator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomPodRotator{}, &CustomPodRotatorList{})
}

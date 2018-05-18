package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApplicationPhase Phase

const (
	ApplicationPhaseNone        = ""
	ApplicationPhasePending     = "Pending"
	ApplicationPhaseInstalling  = "Installing"
	ApplicationPhaseUpgrading   = "Upgrading"
	ApplicationPhaseImplemented = "Implemented"
	ApplicationPhaseFailed      = "Failed"
)

// SDSApplicationList is a list of sds applications.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SDSApplication `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDSApplicationSpec   `json:"spec"`
	Status            SDSApplicationStatus `json:"status"`
}

// SDSApplicationSpec represents a SDS Application - or a helm chart install
// +k8s:openapi-gen=true
type SDSApplicationSpec struct {
	// What tiller should be used
	PackageManager SDSPackageManagerRef `json:"packageManager"`
	// What namespace do we want this application installed in
	Namespace string `json:"namespace"`
	// What is the application release's name
	Name string `json:"name"`
	// Chart Information
	Chart Chart `json:"chart"`
	// What are the values for the Values.yaml file?
	Values string `json:"values"`
}

// Chart represents a SDS Application's Chart information
// +k8s:openapi-gen=true
type Chart struct {
	// What is the chart name
	Name string `json:"chartName"`
	// What is the repository information
	Repository ChartRepository `json:"repository"`
	// What is the chart version
	Version string `json:"version"`
}

// ChartRepository represents a helm chart repository
// +k8s:openapi-gen=true
type ChartRepository struct {
	// What is the repository name
	Name string `json:"name"`
	// What is the URL for the repository?
	URL string `json:"url"`
}

// The status of an application object
// +k8s:openapi-gen=true
type SDSApplicationStatus struct {
	Phase      ApplicationPhase `json:"phase"`
	Ready      bool             `json:"ready"`
	Conditions []Condition      `json:"conditions"`
}

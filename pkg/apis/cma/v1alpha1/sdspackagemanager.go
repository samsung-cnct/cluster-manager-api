package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PackageManagerPhase Phase

const (
	PackageManagerPhaseNone        = ""
	PackageManagerPhasePending     = "Pending"
	PackageManagerPhaseInstalling  = "Installing"
	PackageManagerPhaseUpgrading   = "Upgrading"
	PackageManagerPhaseImplemented = "Implemented"
	PackageManagerPhaseFailed      = "Failed"
)

// SDSPackageManagerList is a list of sds package managers (tiller).
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSPackageManagerList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SDSPackageManager `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSPackageManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDSPackageManagerSpec   `json:"spec"`
	Status            SDSPackageManagerStatus `json:"status"`
}

// SDSPackageManagerRef is a reference to a SDS Package Manager
// +k8s:openapi-gen=true
type SDSPackageManagerRef struct {
	// Name of the SDS Package Manager.  Note that the package manager resource must be in the same namespace right now
	Name string `json:"name"`
}

// SDSPackageManagerSpec represents a SDS Package Manager - or a tiller deployment
// +k8s:openapi-gen=true
type SDSPackageManagerSpec struct {
	// What namespace do we want this package manager installed in
	Namespace string `json:"namespace"`
	// What is the application release's name
	Name string `json:"name"`
	// What version should this package manager be?
	Version string `json:"version"`
	// ServiceAccount to use
	ServiceAccount ServiceAccount `json:"serviceAccount"`
	// Permissions of the Package Manager
	Permissions PackageManagerPermissions `json:"permissions"`
}

// PackageManagerPermissions represents the permissions for the package manager
// +k8s:openapi-gen=true
type PackageManagerPermissions struct {
	// Should this be a cluster wide admin?
	ClusterWide bool `json:"clusterWide"`
	// What namespaces should this package manager administrate on?
	Namespaces []string `json:"namespaces"`
}

// ServiceAccount represents the service account settings
// +k8s:openapi-gen=true
type ServiceAccount struct {
	// Should this be a cluster wide admin?
	Name string `json:"name"`
	// What namespace would we find this service account?
	Namespace string `json:"namespace"`
}

// The status of an package manager object
// +k8s:openapi-gen=true
type SDSPackageManagerStatus struct {
	Phase      PackageManagerPhase `json:"phase"`
	Ready      bool                `json:"ready"`
	Conditions []Condition         `json:"conditions"`
}

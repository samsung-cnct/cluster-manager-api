package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ClusterPhaseNone	= ""
	ClusterPhasePending = "Pending"
	ClusterPhaseWaitingForCluster = "Waiting for cluster"
	ClusterPhaseHaveCluster = "Cluster created"
	ClusterPhaseDeployingPackageManager = "Deploying Package Manager"
	ClusterPhaseHavePackageManager = "Package Manager Installed"
	ClusterPhaseDeployingApplications = "Deploying Applications"
	ClusterPhaseReady = "Ready"
)

// SDSDClusterList is a list of sds clusters.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSClusterList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SDSCluster `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SDSCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDSClusterSpec   `json:"spec"`
	Status            SDSClusterStatus `json:"status"`
}

// SDSClusterSpec represents a SDSCluster spec
// +k8s:openapi-gen=true
type SDSClusterSpec struct {
	// What provider
	Provider	string		`json:"provider",omitempty`
	// What package manager should be used
	PackageManager SDSPackageManagerSpec	`json:"packageManager"`
	// What Charts should be installed
	Applications []SDSApplicationSpec 			`json:"applications,omitempty"`
}

// SDSClusterStatus has the status of the system
// +k8s:openapi-gen=true
type SDSClusterStatus struct {
	Phase      Phase       `json:"phase"`
	Conditions []Condition `json:"conditions"`
	ClusterBuilt bool		`json:clusterBuilt`
	TillerInstalled bool	`json:tillerInstalled`
	AppsInstalled bool		`json:appsInstalled`
}


// SDSClusterRef is a reference to a SDS Cluster
// +k8s:openapi-gen=true
type SDSClusterRef struct {
	// Name of the SDS Cluster.  Note that the cluster must be in the same namespace right now
	Name string `json:"name"`
}

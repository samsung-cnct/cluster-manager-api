package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SDSDClusterList is a list of sds clusters.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SDSDClusterList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SDSCluster `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
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
	// What Charts should be installed
	Charts []Chart 			`json:"chart,omitempty"`
}

// SDSClusterSettings defines the specification of the redis system
// +k8s:openapi-gen=true
type Chart struct {
	Name	  string         `json:"name"`
	Chart     string           `json:"exporter,omitempty"`
	Version   string         `json:"version,omitempty"`
}

// SDSClusterStatus has the status of the system
// +k8s:openapi-gen=true
type SDSClusterStatus struct {
	Phase      Phase       `json:"phase"`
	Conditions []Condition `json:"conditions"`
	Master     string      `json:"master"`
	ClusterBuilt bool		`json:clusterBuilt`
	TillerInstalled bool	`json:tillerInstalled`
	AppsInstalled bool		`json:appsInstalled`
}

// Phase of the RF status
type Phase string

// Condition saves the state information of the redis system
// +k8s:openapi-gen=true
type Condition struct {
	Type           ConditionType `json:"type"`
	Reason         string        `json:"reason"`
	TransitionTime string        `json:"transitionTime"`
}

// ConditionType defines the condition that the redis can have
type ConditionType string

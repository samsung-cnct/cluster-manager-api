package v1alpha1

// Phase of the RF status
type Phase string

// Condition saves the state information of the system
// +k8s:openapi-gen=true
type Condition struct {
	Type           ConditionType `json:"type"`
	Reason         string        `json:"reason"`
	TransitionTime string        `json:"transitionTime"`
}

// ConditionType defines the condition that the condition can have
type ConditionType string

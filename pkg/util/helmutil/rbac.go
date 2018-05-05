package helmutil

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rbac "k8s.io/api/rbac/v1"
)

func GenerateClusterAdminRole(name string) rbac.ClusterRole {
	return rbac.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Rules: []rbac.PolicyRule{
			{
				APIGroups: []string{"*"},
				Resources: []string{"*"},
				Verbs: []string{"*"},
			},
		},
	}
}

func GenerateAdminRole(name string) rbac.Role {
	return rbac.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Rules: []rbac.PolicyRule{
			{
				APIGroups: []string{"*"},
				Resources: []string{"*"},
				Verbs: []string{"*"},
			},
		},
	}
}


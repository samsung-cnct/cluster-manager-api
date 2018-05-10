package k8sutil

import (
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GenerateSingleClusterRolebinding(name string, subject string, subjectNamespace string, roleName string) rbac.ClusterRoleBinding {
	return rbac.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Subjects: []rbac.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      subject,
				Namespace: subjectNamespace,
			},
		},
		RoleRef: rbac.RoleRef{
			Kind:     "ClusterRole",
			Name:     roleName,
			APIGroup: "rbac.authorization.k8s.io",
		},
	}
}

func GenerateSingleRolebinding(name string, subject string, subjectNamespace string, roleName string) rbac.RoleBinding {
	return rbac.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Subjects: []rbac.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      subject,
				Namespace: subjectNamespace,
			},
		},
		RoleRef: rbac.RoleRef{
			Kind:     "Role",
			Name:     roleName,
			APIGroup: "rbac.authorization.k8s.io",
		},
	}
}

func CreateClusterRole(schema rbac.ClusterRole, config *rest.Config) (bool, error) {
	SetLogger()
	if config == nil {
		config = DefaultConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return false, err
	}

	_, err = clientSet.RbacV1().ClusterRoles().Create(&schema)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		logger.Infof("ClusterRole -->%s<-- Cannot be created, error was %v", schema.ObjectMeta.Name, err)
		return false, err
	} else if IsResourceAlreadyExistsError(err) {
		logger.Infof("ClusterRole -->%s<-- Already exists, cannot recreate", schema.ObjectMeta.Name)
		return false, err
	}
	return true, nil
}

func CreateRole(schema rbac.Role, namespace string, config *rest.Config) (bool, error) {
	SetLogger()
	if config == nil {
		config = DefaultConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return false, err
	}

	_, err = clientSet.RbacV1().Roles(namespace).Create(&schema)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		logger.Infof("Role -->%s<-- in namespace -->%s<-- Cannot be created, error was %v", schema.ObjectMeta.Name, namespace, err)
		return false, err
	} else if IsResourceAlreadyExistsError(err) {
		logger.Infof("Role -->%s<-- in namespace -->%s<-- Already exists, cannot recreate", schema.ObjectMeta.Name, namespace)
		return false, err
	}
	return true, nil
}

func CreateClusterRoleBinding(schema rbac.ClusterRoleBinding, config *rest.Config) (bool, error) {
	SetLogger()
	if config == nil {
		config = DefaultConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return false, err
	}

	_, err = clientSet.RbacV1().ClusterRoleBindings().Create(&schema)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		logger.Infof("ClusterRoleBinding -->%s<-- Cannot be created, error was %v", schema.ObjectMeta.Name, err)
		return false, err
	} else if IsResourceAlreadyExistsError(err) {
		logger.Infof("ClusterRoleBinding -->%s<-- Already exists, cannot recreate", schema.ObjectMeta.Name)
		return false, err
	}
	return true, nil
}

func CreateRoleBinding(schema rbac.RoleBinding, namespace string, config *rest.Config) (bool, error) {
	SetLogger()
	if config == nil {
		config = DefaultConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return false, err
	}

	_, err = clientSet.RbacV1().RoleBindings(namespace).Create(&schema)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		logger.Infof("RoleBinding -->%s<-- in namespace -->%s<-- Cannot be created, error was %v", schema.ObjectMeta.Name, namespace, err)
		return false, err
	} else if IsResourceAlreadyExistsError(err) {
		logger.Infof("RoleBinding -->%s<-- in namespace -->%s<-- Already exists, cannot recreate", schema.ObjectMeta.Name, namespace)
		return false, err
	}
	return true, nil
}

package k8sutil

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GenerateNamespace(name string) corev1.Namespace {
	return corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}

func CreateNamespace(schema corev1.Namespace, config *rest.Config) (bool, error) {
	SetLogger()
	if config == nil {
		config = DefaultConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return false, err
	}

	_, err = clientSet.CoreV1().Namespaces().Create(&schema)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		logger.Infof("Namespace -->%s<--  Cannot be created, error was %v", schema.ObjectMeta.Name, err)
		return false, err
	} else if IsResourceAlreadyExistsError(err) {
		logger.Infof("Namespace -->%s<-- Already exists, cannot recreate", schema.ObjectMeta.Name)
		return false, err
	}
	return true, nil
}

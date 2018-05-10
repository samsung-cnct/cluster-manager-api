package k8sutil

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GenerateServiceAccount(name string) corev1.ServiceAccount {
	return corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}

func CreateServiceAccount(schema corev1.ServiceAccount, namespace string, config *rest.Config) (bool, error) {
	SetLogger()
	if config == nil {
		config = DefaultConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return false, err
	}

	_, err = clientSet.CoreV1().ServiceAccounts(namespace).Create(&schema)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		logger.Infof("ServiceAccount -->%s<-- in namespace -->%s<-- Cannot be created, error was %v", schema.ObjectMeta.Name, namespace, err)
		return false, err
	} else if IsResourceAlreadyExistsError(err) {
		logger.Infof("ServiceAccount -->%s<-- in namespace -->%s<-- Already exists, cannot recreate", schema.ObjectMeta.Name, namespace)
		return false, err
	}
	return true, nil
}

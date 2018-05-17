package k8sutil

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GenerateJob(name string, jobSpec batchv1.JobSpec) batchv1.Job {
	return batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: jobSpec,
	}
}

func CreateJob(schema batchv1.Job, namespace string, config *rest.Config) (bool, error) {
	SetLogger()
	if config == nil {
		config = DefaultConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("Cannot establish a client connection to kubernetes: %v", err)
		return false, err
	}

	_, err = clientSet.BatchV1().Jobs(namespace).Create(&schema)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		logger.Infof("Job -->%s<-- in namespace -->%s<-- Cannot be created, error was %v", schema.ObjectMeta.Name, "default", err)
		return false, err
	} else if IsResourceAlreadyExistsError(err) {
		logger.Infof("Job -->%s<-- in namespace -->%s<-- Already exists, cannot recreate", schema.ObjectMeta.Name, namespace)
		return false, err
	}
	return true, nil
}

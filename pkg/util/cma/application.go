package cma

import (
	sdsapi "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type SDSApplicationOptions struct {
	Name           string
	Namespace      string
	Values         string
	PackageManager string
	Chart          Chart
}

type Chart struct {
	Name       string
	Repository ChartRepository
	Version    string
}

type ChartRepository struct {
	Name string
	URL  string
}

func GenerateSDSApplication(options SDSApplicationOptions) sdsapi.SDSApplication {
	return sdsapi.SDSApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name,
		},
		Spec: sdsapi.SDSApplicationSpec{
			Name:           options.Name,
			Namespace:      options.Namespace,
			Values:         options.Values,
			PackageManager: sdsapi.SDSPackageManagerRef{Name: options.PackageManager},
			Chart: sdsapi.Chart{
				Name:    options.Chart.Name,
				Version: options.Chart.Version,
				Repository: sdsapi.ChartRepository{
					Name: options.Chart.Repository.Name,
					URL:  options.Chart.Repository.URL,
				},
			},
		},
	}
}

func CreateSDSApplication(application sdsapi.SDSApplication, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	_, err = client.CmaV1alpha1().SDSApplications(namespace).Create(&application)
	if err != nil && !k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("Application -->%s<-- Cannot be created, error was %v", application.ObjectMeta.Name, err)
		return false, err
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("Application -->%s<-- Already exists, cannot recreate", application.ObjectMeta.Name)
		return false, err
	}

	return true, nil
}

func UpdateSDSApplication(application sdsapi.SDSApplication, namespace string, config *rest.Config) (*sdsapi.SDSApplication, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	updatedApplication, err := client.CmaV1alpha1().SDSApplications(namespace).Update(&application)
	if err != nil {
		logger.Infof("Application -->%s<-- Cannot be updated, error was %v", application.ObjectMeta.Name, err)
		return nil, err
	}

	return updatedApplication, nil
}

func GetSDSApplication(name string, namespace string, config *rest.Config) (*sdsapi.SDSApplication, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	application, err := client.CmaV1alpha1().SDSApplications(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		logger.Infof("Application -->%s<-- Cannot be fetched, error was %v", name, err)
		return nil, err
	}

	return application, nil
}

func DeleteSDSApplication(name string, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	err = client.CmaV1alpha1().SDSApplications(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		logger.Infof("Application -->%s<-- Cannot be deleted, error was %v", name, err)
		return false, err
	}

	return true, nil
}

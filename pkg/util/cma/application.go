package cma

import (
	sdsapi "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// Options to for use with the GenerateSDSApplication method
type SDSApplicationOptions struct {
	// Name of the application (what will show up in helm list)
	Name string
	// Namespace of the application within the target kubernetes environment
	Namespace string
	// Values in the standard helm --values command line argument.  This will be written to disk and used
	Values string
	// What package manager (tiller) should be used when installing this application
	PackageManager string
	// The chart information (name, repository, version, etc.)
	Chart Chart
}

// Chart represents the Chart part of the SDSApplicationOptions struct
type Chart struct {
	// Name of the Chart - should be repository/chart-name like fooinc/my-chart
	Name string
	// TODO Make ChartRepository an array
	Repository ChartRepository
	// Version of the chart to install
	Version string
}

// Representation of a Chart Repository - used in Chart which is then used by SDSApplicationOptions which is used by
// GenerateSDSApplication
type ChartRepository struct {
	// Name that will be used within the helm command.  So if it is called fooinc, chart name would be fooinc/my-chart
	Name string
	// The URL to the chart repository
	URL string
}

// This generic helper method will create a SDSApplication object for us with the CMA API
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

// This method will create an SDS Application in the cluster provided by the config value.
// The namespace value will be what namespace the custom resource (SDSApplication) will be installed in, not where the
// eventual application will be installed in.
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

// This method will update an existing SDSApplication resource in the cluster provided by the config value.
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

// This method will return an existing SDSApplication resource in the cluster provided by the config value.
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

// This method will delete an existing SDSApplication resource in the cluster provided by the config value.
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

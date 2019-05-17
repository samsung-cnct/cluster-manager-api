package cmak8sutil

import (
	"github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeSchema "k8s.io/apimachinery/pkg/runtime/schema"
)

func (c *Client) CreateApplication(name string, packageManager string, application Application) error {
	adjustedName := c.getAdjustedApplicationName(name, packageManager, application.Cluster)

	annotations := make(map[string]string)
	annotations[CallbackURLAnnotation] = application.CallbackURL
	annotations[RequestIDAnnotation] = application.RequestID

	sdsCluster, clusterErr := c.getClusterRaw(application.Cluster)
	if clusterErr != nil {
		return clusterErr
	}

	ownerRefs := []v1.OwnerReference{
		*v1.NewControllerRef(sdsCluster,
			runtimeSchema.GroupVersionKind{
				Group:   v1alpha1.SchemeGroupVersion.Group,
				Version: v1alpha1.SchemeGroupVersion.Version,
				Kind:    "SDSCluster",
			}),
	}

	_, err := c.applicationClient.Create(&v1alpha1.SDSApplication{
		ObjectMeta: v1.ObjectMeta{Name: adjustedName, Annotations: annotations, OwnerReferences: ownerRefs},
		Spec: v1alpha1.SDSApplicationSpec{
			PackageManager: v1alpha1.SDSPackageManagerRef{
				Name: application.PackageManager,
			},
			Name:      name,
			Namespace: application.Namespace,
			Chart: v1alpha1.Chart{
				Name: application.Chart.Name,
				Repository: v1alpha1.ChartRepository{
					Name: application.Chart.Repository.Name,
					URL:  application.Chart.Repository.URL,
				},
				ChartPayload: application.Chart.ChartPayload,
				Version:      application.Chart.Version,
			},
			Values: application.Values,
			Cluster: v1alpha1.SDSClusterRef{
				Name: application.Cluster,
			},
		},
	})
	return err
}

func (c *Client) getApplicationRaw(name string) (*v1alpha1.SDSApplication, error) {
	return c.applicationClient.Get(name, v1.GetOptions{})
}

func (c *Client) GetApplication(name string, packageManager string, clusterName string) (Application, error) {
	result, err := c.getApplicationRaw(c.getAdjustedApplicationName(name, packageManager, clusterName))
	if err != nil {
		return Application{}, nil
	}
	return Application{
		CallbackURL: result.Annotations[CallbackURLAnnotation],
		Cluster:     result.Spec.Cluster.Name,
		Chart: Chart{
			Name: result.Spec.Chart.Name,
			Repository: ChartRepository{
				Name: result.Spec.Chart.Repository.Name,
				URL:  result.Spec.Chart.Repository.URL,
			},
			ChartPayload: result.Spec.Chart.ChartPayload,
			Version:      result.Spec.Chart.Version,
		},
		Namespace:      result.Spec.Namespace,
		PackageManager: result.Spec.PackageManager.Name,
		RequestID:      result.Annotations[RequestIDAnnotation],
		Values:         result.Spec.Values,
	}, nil
}

func (c *Client) UpdateOrCreateApplication(name string, packageManager string, application Application) error {
	result, err := c.getApplicationRaw(c.getAdjustedApplicationName(name, packageManager, application.Cluster))
	if err != nil {
		// Let's assume there is no application, so we create it
		return c.CreateApplication(name, packageManager, application)
	}
	sdsCluster, clusterErr := c.getClusterRaw(application.Cluster)
	if clusterErr != nil {
		return clusterErr
	}

	ownerRefs := []v1.OwnerReference{
		*v1.NewControllerRef(sdsCluster,
			runtimeSchema.GroupVersionKind{
				Group:   v1alpha1.SchemeGroupVersion.Group,
				Version: v1alpha1.SchemeGroupVersion.Version,
				Kind:    "SDSCluster",
			}),
	}

	result.Annotations[CallbackURLAnnotation] = application.CallbackURL
	result.Annotations[RequestIDAnnotation] = application.RequestID
	result.OwnerReferences = ownerRefs
	result.Spec = v1alpha1.SDSApplicationSpec{
		PackageManager: v1alpha1.SDSPackageManagerRef{
			Name: application.PackageManager,
		},
		Name: name,
		Chart: v1alpha1.Chart{
			Name: application.Chart.Name,
			Repository: v1alpha1.ChartRepository{
				Name: application.Chart.Repository.Name,
				URL:  application.Chart.Repository.URL,
			},
			ChartPayload: application.Chart.ChartPayload,
			Version:      application.Chart.Version,
		},
		Values: application.Values,
		Cluster: v1alpha1.SDSClusterRef{
			Name: application.Cluster,
		},
	}

	_, err = c.applicationClient.Update(result)
	return err
}

func (c *Client) DeleteApplication(name string, packageManager string, clusterName string) error {
	return c.applicationClient.Delete(c.getAdjustedApplicationName(name, packageManager, clusterName), nil)
}

func (c *Client) ChangeApplicationStatus(name string, packageManager string, clusterName string, status string) error {
	result, err := c.getApplicationRaw(c.getAdjustedApplicationName(name, packageManager, clusterName))
	if err != nil {
		return err
	}
	result.Status.Phase = v1alpha1.ApplicationPhase(status)
	_, err = c.applicationClient.Update(result)
	return err
}

package cmak8sutil

import (
	"github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeSchema "k8s.io/apimachinery/pkg/runtime/schema"
)

func (c *Client) CreatePackageManager(name string, packageManager PackageManager) error {
	adjustedName := c.getAdjustedName(name, packageManager.Cluster)

	annotations := make(map[string]string)
	annotations[CallbackURLAnnotation] = packageManager.CallbackURL
	annotations[RequestIDAnnotation] = packageManager.RequestID

	sdsCluster, clusterErr := c.getClusterRaw(packageManager.Cluster)
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

	_, err := c.packageManagerClient.Create(&v1alpha1.SDSPackageManager{
		ObjectMeta: v1.ObjectMeta{Name: adjustedName, Annotations: annotations, OwnerReferences: ownerRefs},
		Spec: v1alpha1.SDSPackageManagerSpec{
			Name:      name,
			Namespace: packageManager.Namespace,
			Version:   packageManager.Version,
			Image:     packageManager.Image,
			ServiceAccount: v1alpha1.ServiceAccount{
				Name:      name,
				Namespace: packageManager.Namespace,
			},
			Permissions: v1alpha1.PackageManagerPermissions{
				ClusterWide: packageManager.ClusterWide,
				Namespaces:  packageManager.AdminNamespaces,
			},
			Cluster: v1alpha1.SDSClusterRef{
				Name: packageManager.Cluster,
			},
		},
	})
	return err
}

func (c *Client) getPackageManagerRaw(name string) (*v1alpha1.SDSPackageManager, error) {
	return c.packageManagerClient.Get(name, v1.GetOptions{})
}

func (c *Client) GetPackageManager(name string, clusterName string) (PackageManager, error) {
	result, err := c.getPackageManagerRaw(c.getAdjustedName(name, clusterName))
	if err != nil {
		return PackageManager{}, nil
	}
	return PackageManager{
		AdminNamespaces: result.Spec.Permissions.Namespaces,
		CallbackURL:     result.Annotations[CallbackURLAnnotation],
		Cluster:         result.Spec.Cluster.Name,
		ClusterWide:     result.Spec.Permissions.ClusterWide,
		Image:           result.Spec.Image,
		Name:            result.Name,
		Namespace:       result.Namespace,
		RequestID:       result.Annotations[RequestIDAnnotation],
		Version:         result.Spec.Version,
	}, nil
}

func (c *Client) UpdateOrCreatePackageManager(name string, packageManager PackageManager) error {
	result, err := c.getPackageManagerRaw(c.getAdjustedName(name, packageManager.Cluster))
	if err != nil {
		// Let's assume there is no application, so we create it
		return c.CreatePackageManager(name, packageManager)
	}

	sdsCluster, clusterErr := c.getClusterRaw(packageManager.Cluster)
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

	result.Annotations[CallbackURLAnnotation] = packageManager.CallbackURL
	result.Annotations[RequestIDAnnotation] = packageManager.RequestID
	result.OwnerReferences = ownerRefs
	result.Spec = v1alpha1.SDSPackageManagerSpec{
		Name:      name,
		Namespace: packageManager.Namespace,
		Version:   packageManager.Version,
		Image:     packageManager.Image,
		ServiceAccount: v1alpha1.ServiceAccount{
			Name:      name,
			Namespace: packageManager.Namespace,
		},
		Permissions: v1alpha1.PackageManagerPermissions{
			ClusterWide: packageManager.ClusterWide,
			Namespaces:  packageManager.AdminNamespaces,
		},
		Cluster: v1alpha1.SDSClusterRef{
			Name: packageManager.Cluster,
		},
	}

	_, err = c.packageManagerClient.Update(result)
	return err
}

func (c *Client) DeletePackageManager(name string, clusterName string) error {
	return c.packageManagerClient.Delete(c.getAdjustedName(name, clusterName), nil)
}

func (c *Client) ChangePackageManagerStatus(name string, clusterName string, status string) error {
	result, err := c.getPackageManagerRaw(c.getAdjustedName(name, clusterName))
	if err != nil {
		return err
	}
	result.Status.Phase = v1alpha1.PackageManagerPhase(status)
	_, err = c.packageManagerClient.Update(result)
	return err
}

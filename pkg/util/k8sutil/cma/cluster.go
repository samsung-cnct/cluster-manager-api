package cmak8sutil

import (
	"github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateCluster(name string, cluster Cluster) error {
	annotations := make(map[string]string)
	annotations[CallbackURLAnnotation] = cluster.CallbackURL
	annotations[RequestIDAnnotation] = cluster.RequestID

	_, err := c.clusterClient.Create(&v1alpha1.SDSCluster{
		ObjectMeta: v1.ObjectMeta{Name: name, Annotations: annotations},
		Spec: v1alpha1.SDSClusterSpec{
			Provider: cluster.Provider,
		},
	})
	return err
}

func (c *Client) getClusterRaw(name string) (*v1alpha1.SDSCluster, error) {
	return c.clusterClient.Get(name, v1.GetOptions{})
}

func (c *Client) GetCluster(name string) (Cluster, error) {
	result, err := c.clusterClient.Get(name, v1.GetOptions{})
	if err != nil {
		return Cluster{}, nil
	}
	return Cluster{
		CallbackURL: result.Annotations[CallbackURLAnnotation],
		Provider:    result.Spec.Provider,
		RequestID:   result.Annotations[RequestIDAnnotation],
	}, nil
}

func (c *Client) UpdateOrCreateCluster(name string, cluster Cluster) error {
	result, err := c.getClusterRaw(name)
	if err != nil {
		// Let's assume there is no application, so we create it
		return c.CreateCluster(name, cluster)
	}

	result.Annotations[CallbackURLAnnotation] = cluster.CallbackURL
	result.Annotations[RequestIDAnnotation] = cluster.RequestID
	result.Spec.Provider = cluster.Provider

	_, err = c.clusterClient.Update(result)
	return err

}

func (c *Client) DeleteCluster(name string) error {
	return c.clusterClient.Delete(name, nil)
}

func (c *Client) ChangeClusterStatus(name string, status string) error {
	result, err := c.getClusterRaw(name)
	if err != nil {
		return err
	}
	result.Status.Phase = v1alpha1.Phase(status)
	_, err = c.clusterClient.Update(result)
	return err
}

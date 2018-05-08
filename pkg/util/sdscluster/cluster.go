package sdscluster

import (
	sdsapi "github.com/samsung-cnct/cluster-manager-api/pkg/apis/sdscluster/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/client-go/rest"
)

type SDSClusterOptions struct {
	Name string
	Provider string
}

func GenerateSDSCluster(options SDSClusterOptions) sdsapi.SDSCluster {
	return sdsapi.SDSCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name,
		},
		Spec: sdsapi.SDSClusterSpec{
			Provider: options.Provider,
		},
	}
}

func CreateKrakenCluster(cluster sdsapi.SDSCluster, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	_, err = client.SdsclusterV1alpha1().SDSClusters(namespace).Create(&cluster)
	if err != nil && !k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("KrakenCluster -->%s<-- Cannot be created, error was %v", cluster.ObjectMeta.Name, err)
		return false, err
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("KrakenCluster -->%s<-- Already exists, cannot recreate", cluster.ObjectMeta.Name)
		return false, err
	}

	return true, err
}

func DeleteKrakenCluster(name string, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	err = client.SdsclusterV1alpha1().SDSClusters(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		logger.Infof("KrakenCluster -->%s<-- Cannot be Deleted, error was %v", name, err)
		return false, err
	}
	return true, err
}

func GetKrakenCluster(name string, namespace string, config *rest.Config) (*sdsapi.SDSCluster, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	cluster, err := client.SdsclusterV1alpha1().SDSClusters(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		logger.Infof("KrakenCluster -->%s<-- Cannot be Retrieved, error was %v", name, err)
		return nil, err
	}
	return cluster, err
}
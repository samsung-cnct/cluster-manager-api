package azure

import (
	"fmt"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaks"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/azure"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
	"github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	"github.com/sirupsen/logrus"
)

func (c *Client) UpdateCredentials(in *pb.UpdateAzureCredentialsMsg) (*pb.UpdateAzureCredentialsReply, error) {
	err := c.secretClient.UpdateOrCreateCredentials(in.Name, azurek8sutil.Credentials{
		AppID:          in.Credentials.AppId,
		Tenant:         in.Credentials.Tenant,
		Password:       in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return &pb.UpdateAzureCredentialsReply{}, err
	}
	return &pb.UpdateAzureCredentialsReply{Ok: true}, nil

}

func (c *Client) CreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	var instanceGroups []cmaaks.InstanceGroup
	for _, j := range in.Provider.GetAzure().InstanceGroups {
		instanceGroups = append(instanceGroups, cmaaks.InstanceGroup{
			Name:        j.Name,
			Type:        j.Type,
			MinQuantity: int(j.MinQuantity),
			MaxQuantity: int(j.MaxQuantity),
		})
	}
	result, err := c.cmaAKSClient.CreateCluster(cmaaks.CreateClusterInput{
		Name:       in.Name,
		K8SVersion: in.Provider.K8SVersion,
		Azure: cmaaks.AzureSpec{
			Location: in.Provider.GetAzure().Location,
			Credentials: cmaaks.Credentials{
				AppID:          in.Provider.GetAzure().Credentials.AppId,
				Tenant:         in.Provider.GetAzure().Credentials.Tenant,
				Password:       in.Provider.GetAzure().Credentials.Password,
				SubscriptionID: in.Provider.GetAzure().Credentials.SubscriptionId,
			},
			ClusterServiceAccount: cmaaks.ClusterServiceAccount{
				ClientID:     in.Provider.GetAzure().Credentials.AppId,
				ClientSecret: in.Provider.GetAzure().Credentials.Password,
			},
			InstanceGroups: instanceGroups,
		},
		HighAvailability: in.Provider.HighAvailability,
		NetworkFabric:    in.Provider.NetworkFabric,
	})
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}

	// Cluster Creation was successful, going to save the credentials
	err = c.secretClient.CreateCredentials(in.Name, azurek8sutil.Credentials{
		AppID:          in.Provider.GetAzure().Credentials.AppId,
		Tenant:         in.Provider.GetAzure().Credentials.Tenant,
		Password:       in.Provider.GetAzure().Credentials.Password,
		SubscriptionID: in.Provider.GetAzure().Credentials.SubscriptionId,
	})

	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set AKS credentials into kubernetes, this is bad")
	}

	// Now going to create K8S CR
	err = c.cmaK8sClient.CreateCluster(in.Name, cmak8sutil.Cluster{
		CallbackURL: in.Callback.Url,
		Provider:    "azure",
		RequestID:   in.Callback.RequestId,
	})

	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set AKS credentials into kubernetes, this is bad")
	}

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     result.Cluster.ID,
			Name:   result.Cluster.Name,
			Status: pb.ClusterStatus_PROVISIONING,
		},
	}, nil
}

func (c *Client) GetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	credentials, updateCache, err := c.reconcileCredentials(in.Name, in.Azure)
	if err != nil {
		return &pb.GetClusterReply{}, err
	}

	result, err := c.cmaAKSClient.GetCluster(cmaaks.GetClusterInput{
		Name: in.Name,
		Credentials: cmaaks.Credentials{
			AppID:          credentials.AppID,
			Tenant:         credentials.Tenant,
			Password:       credentials.Password,
			SubscriptionID: credentials.SubscriptionID,
		},
	})
	if err != nil {
		return &pb.GetClusterReply{}, err
	}

	// Processing output
	if updateCache {
		err = c.updateCachedCredentials(in.Name, credentials)
		if err != nil {
			// Could not update the credentials, let's log that
			logrus.Errorf("could not update credentials for cluster -->%s<--, error was %s", in.Name, err)
		}
	}

	enumeratedStatus, found := pb.ClusterStatus_value[result.Cluster.Status]
	if !found {
		enumeratedStatus = 0
	}

	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         result.Cluster.ID,
			Name:       result.Cluster.Name,
			Status:     pb.ClusterStatus(enumeratedStatus),
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}, nil
}

func (c *Client) GetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error) {
	var clusters []*pb.ClusterItem
	result, err := c.cmaAKSClient.ListClusters(cmaaks.ListClusterInput{
		Credentials: cmaaks.Credentials{
			AppID:          in.GetAzure().AppId,
			Tenant:         in.GetAzure().Tenant,
			Password:       in.GetAzure().Password,
			SubscriptionID: in.GetAzure().SubscriptionId,
		},
	})
	if err != nil {
		return &pb.GetClusterListReply{}, err
	}
	for _, j := range result.Clusters {
		enumeratedStatus, found := pb.ClusterStatus_value[j.Status]
		if !found {
			enumeratedStatus = 0
		}
		clusters = append(clusters, &pb.ClusterItem{
			Id:     j.ID,
			Name:   j.Name,
			Status: pb.ClusterStatus(enumeratedStatus),
		})
	}
	return &pb.GetClusterListReply{
		Ok:       true,
		Clusters: clusters,
	}, nil
}

func (c *Client) DeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	credentials, _, err := c.reconcileCredentials(in.Name, in.Azure)
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	result, err := c.cmaAKSClient.DeleteCluster(cmaaks.DeleteClusterInput{
		Name: in.Name,
		Credentials: cmaaks.Credentials{
			AppID:          credentials.AppID,
			Tenant:         credentials.Tenant,
			Password:       credentials.Password,
			SubscriptionID: credentials.SubscriptionID,
		},
	})
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}

	// Now going to create K8S CR
	err = c.cmaK8sClient.UpdateOrCreateCluster(in.Name, cmak8sutil.Cluster{
		CallbackURL: in.Callback.Url,
		Provider:    in.Provider.String(),
		RequestID:   in.Callback.RequestId,
	})
	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set AKS credentials into kubernetes, this is bad")
	}
	err = c.cmaK8sClient.ChangeClusterStatus(in.Name, v1alpha1.ClusterPhaseDeleting)
	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set AKS credentials into kubernetes, this is bad")
	}

	// Deleting credentials
	err = c.secretClient.DeleteCredentials(in.Name)
	if err != nil {
		// Could not delete the credentials, let's log that
		logrus.Errorf("could not delete credentials for cluster -->%s<--, error was %s", in.Name, err)
	}

	return &pb.DeleteClusterReply{
		Ok:     true,
		Status: result.Status,
	}, nil
}

func (c *Client) GetClusterUpgrades(in *pb.GetUpgradeClusterInformationMsg) (output *pb.GetUpgradeClusterInformationReply, err error) {
	output = &pb.GetUpgradeClusterInformationReply{}
	credentials, updateCache, err := c.reconcileCredentials(in.Name, in.Azure)
	if err != nil {
		return
	}

	result, err := c.cmaAKSClient.GetClusterUpgrades(cmaaks.GetClusterUpgradesInput{
		Name: in.Name,
		Credentials: cmaaks.Credentials{
			AppID:          credentials.AppID,
			Tenant:         credentials.Tenant,
			Password:       credentials.Password,
			SubscriptionID: credentials.SubscriptionID,
		},
	})
	if err != nil {
		return
	}

	// Processing output
	if updateCache {
		updateErr := c.updateCachedCredentials(in.Name, credentials)
		if updateErr != nil {
			// Could not update the credentials, let's log that
			logrus.Errorf("could not update credentials for cluster -->%s<--, error was %s", in.Name, updateErr)
		}
	}

	for _, j := range result.Versions {
		output.Versions = append(output.Versions, j)
	}

	return
}

func (c *Client) ClusterUpgrade(in *pb.UpgradeClusterMsg) (output *pb.UpgradeClusterReply, err error) {
	output = &pb.UpgradeClusterReply{}
	credentials, updateCache, err := c.reconcileCredentials(in.Name, in.Azure)
	if err != nil {
		return
	}

	_, err = c.cmaAKSClient.ClusterUpgrade(cmaaks.ClusterUpgradeInput{
		Name:    in.Name,
		Version: in.Version,
		Credentials: cmaaks.Credentials{
			AppID:          credentials.AppID,
			Tenant:         credentials.Tenant,
			Password:       credentials.Password,
			SubscriptionID: credentials.SubscriptionID,
		},
	})
	if err != nil {
		return
	}

	// Processing output
	if updateCache {
		updateErr := c.updateCachedCredentials(in.Name, credentials)
		if updateErr != nil {
			// Could not update the credentials, let's log that
			logrus.Errorf("could not update credentials for cluster -->%s<--, error was %s", in.Name, updateErr)
		}
	}

	// Now going to create K8S CR
	err = c.cmaK8sClient.UpdateOrCreateCluster(in.Name, cmak8sutil.Cluster{
		CallbackURL: in.Callback.Url,
		Provider:    in.Provider.String(),
		RequestID:   in.Callback.RequestId,
	})
	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set AKS credentials into kubernetes, this is bad")
	}
	err = c.cmaK8sClient.ChangeClusterStatus(in.Name, v1alpha1.ClusterPhaseUpgrading)
	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set AKS credentials into kubernetes, this is bad")
	}

	output.Ok = true
	return
}

func (c *Client) reconcileCredentials(clusterName string, providedCredentials *pb.AzureCredentials) (credentials azurek8sutil.Credentials, updateCache bool, err error) {
	logrus.Errorf("Reconciling credentials")
	if providedCredentials != nil &&
		providedCredentials.AppId != "" &&
		providedCredentials.Tenant != "" &&
		providedCredentials.Password != "" &&
		providedCredentials.SubscriptionId != "" {
		logrus.Errorf("Using provided credentials")
		return azurek8sutil.Credentials{
			AppID:          providedCredentials.AppId,
			Tenant:         providedCredentials.Tenant,
			Password:       providedCredentials.Password,
			SubscriptionID: providedCredentials.SubscriptionId,
		}, true, nil
	}
	cacheResult, err := c.secretClient.GetCredentials(clusterName)
	if err != nil {
		return azurek8sutil.Credentials{}, false, err
	}
	logrus.Errorf("Using cached credentials")
	return cacheResult, false, nil
}

func (c *Client) updateCachedCredentials(clusterName string, credentials azurek8sutil.Credentials) (err error) {
	logrus.Errorf("Updating cached credentials")
	return c.secretClient.UpdateOrCreateCredentials(clusterName, credentials)
}

func (c *Client) AdjustCluster(in *pb.AdjustClusterMsg) (output *pb.AdjustClusterReply, err error) {
	output = &pb.AdjustClusterReply{}
	credentials, updateCache, err := c.reconcileCredentials(in.Name, in.GetAzure().Credentials)
	if err != nil {
		return
	}

	// find out the current number of nodes in the provided node pool
	currentNodePool, err := c.cmaAKSClient.GetClusterNodeCount(cmaaks.GetClusterNodeCountInput{
		Name: in.Name,
		Credentials: cmaaks.Credentials{
			AppID:          credentials.AppID,
			Tenant:         credentials.Tenant,
			Password:       credentials.Password,
			SubscriptionID: credentials.SubscriptionID,
		},
	})
	if err != nil {
		return
	}

	var newCount int32
	// check if adding
	if in.GetAzure().AddCount > 0 {
		newCount = currentNodePool.Count + in.GetAzure().AddCount
		// check if removing
	} else if in.GetAzure().RemoveCount > 0 {
		if currentNodePool.Count-in.GetAzure().RemoveCount < 1 {
			return output, fmt.Errorf("can not scale below 1 node on cluster -->%s<---", in.Name)
		}
		newCount = currentNodePool.Count - in.GetAzure().RemoveCount
	}

	// check that we are increasing or decreasing the number of nodes if not change then use existing count
	if newCount == 0 {
		newCount = currentNodePool.Count
	}

	_, err = c.cmaAKSClient.ScaleCluster(cmaaks.ScaleClusterInput{
		Name: in.Name,
		Credentials: cmaaks.Credentials{
			AppID:          credentials.AppID,
			Tenant:         credentials.Tenant,
			Password:       credentials.Password,
			SubscriptionID: credentials.SubscriptionID,
		},
		NodePool: in.GetAzure().NodePool,
		Count:    newCount,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to scaling cluster -->%s<--: error was %s", in.Name, err)
	}

	// Processing output
	if updateCache {
		updateErr := c.updateCachedCredentials(in.Name, credentials)
		if updateErr != nil {
			// Could not update the credentials, let's log that
			logrus.Errorf("could not update credentials for cluster -->%s<--, error was %s", in.Name, updateErr)
		}
	}

	// Now going to create K8S CR
	err = c.cmaK8sClient.UpdateOrCreateCluster(in.Name, cmak8sutil.Cluster{
		CallbackURL: in.Callback.Url,
		Provider:    in.Provider.String(),
		RequestID:   in.Callback.RequestId,
	})
	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set Cluster CR into kubernetes, this is bad")
	}
	err = c.cmaK8sClient.ChangeClusterStatus(in.Name, v1alpha1.ClusterPhaseUpgrading)
	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set Cluster CR into kubernetes, this is bad")
	}

	if err != nil {
		return &pb.AdjustClusterReply{}, err
	}

	return &pb.AdjustClusterReply{
		Ok: true,
	}, nil
}

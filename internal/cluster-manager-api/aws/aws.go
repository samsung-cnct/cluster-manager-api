package aws

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaws"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/aws"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
	"github.com/samsung-cnct/cma-operator/pkg/apis/cma/v1alpha1"
	"github.com/sirupsen/logrus"
)

func (c *Client) UpdateCredentials(in *pb.UpdateAWSCredentialsMsg) (*pb.UpdateAWSCredentialsReply, error) {
	err := c.secretClient.UpdateOrCreateCredentials(in.Name, awsk8sutil.Credentials{
		Region:          in.Credentials.Region,
		SecretKeyID:     in.Credentials.SecretKeyId,
		SecretAccessKey: in.Credentials.SecretAccessKey,
	})
	if err != nil {
		return &pb.UpdateAWSCredentialsReply{}, err
	}
	return &pb.UpdateAWSCredentialsReply{Ok: true}, nil
}

func (c *Client) CreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	err := c.validateCreateClusterInput(in)
	if err != nil {
		return nil, err
	}
	var instanceGroups []cmaaws.InstanceGroup
	for _, j := range in.Provider.GetAws().InstanceGroups {
		instanceGroups = append(instanceGroups, cmaaws.InstanceGroup{
			Type:        j.Type,
			MinQuantity: int(j.MinQuantity),
			MaxQuantity: int(j.MaxQuantity),
		})
	}
	result, err := c.cmaAWSClient.CreateCluster(cmaaws.CreateClusterInput{
		Name:       in.Name,
		K8SVersion: in.Provider.K8SVersion,
		AWS: cmaaws.AWSSpec{
			DataCenter: cmaaws.DataCenter{
				Region:            in.Provider.GetAws().DataCenter.Region,
				AvailabilityZones: in.Provider.GetAws().DataCenter.AvailabilityZones,
			},
			Credentials: cmaaws.Credentials{
				Region:          in.Provider.GetAws().Credentials.Region,
				SecretKeyID:     in.Provider.GetAws().Credentials.SecretKeyId,
				SecretAccessKey: in.Provider.GetAws().Credentials.SecretAccessKey,
			},
			PreconfiguredItems: cmaaws.PreconfiguredItems{
				VPCID:           in.Provider.GetAws().Resources.VpcId,
				SecurityGroupID: in.Provider.GetAws().Resources.SecurityGroupId,
				IAMRoleARN:      in.Provider.GetAws().Resources.IamRoleArn,
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
	err = c.secretClient.CreateCredentials(in.Name, awsk8sutil.Credentials{
		Region:          in.Provider.GetAws().Credentials.Region,
		SecretKeyID:     in.Provider.GetAws().Credentials.SecretKeyId,
		SecretAccessKey: in.Provider.GetAws().Credentials.SecretAccessKey,
	})

	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set AWS credentials into kubernetes, this is bad")
	}

	// Now going to create K8S CR
	err = c.cmaK8sClient.CreateCluster(in.Name, cmak8sutil.Cluster{
		CallbackURL: in.Callback.Url,
		Provider:    "aws",
		RequestID:   in.Callback.RequestId,
	})

	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentials to kubernetes
		// TODO Going to log for now
		logrus.Errorf("Could not set Cluster CR into kubernetes, this is bad")
	}

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     result.Cluster.ID,
			Name:   result.Cluster.Name,
			Status: pb.ClusterStatus(pb.ClusterStatus_value[result.Cluster.Status]),
		},
	}, nil
}

func (c *Client) GetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	credentials, updateCache, err := c.reconcileCredentials(in.Name, in.Aws)
	if err != nil {
		return &pb.GetClusterReply{}, err
	}
	result, err := c.cmaAWSClient.GetCluster(cmaaws.GetClusterInput{
		Name: in.Name,
		Credentials: cmaaws.Credentials{
			Region:          credentials.Region,
			SecretKeyID:     credentials.SecretKeyID,
			SecretAccessKey: credentials.SecretAccessKey,
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

	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         result.Cluster.ID,
			Name:       result.Cluster.Name,
			Status:     pb.ClusterStatus(pb.ClusterStatus_value[result.Cluster.Status]),
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}, nil
}

func (c *Client) GetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error) {
	var clusters []*pb.ClusterItem
	result, err := c.cmaAWSClient.ListClusters(cmaaws.ListClusterInput{
		Credentials: cmaaws.Credentials{
			Region:          in.GetAws().Region,
			SecretKeyID:     in.GetAws().SecretKeyId,
			SecretAccessKey: in.GetAws().SecretAccessKey,
		},
	})
	if err != nil {
		return &pb.GetClusterListReply{}, err
	}
	for _, j := range result.Clusters {
		clusters = append(clusters, &pb.ClusterItem{
			Id:     j.ID,
			Name:   j.Name,
			Status: pb.ClusterStatus(pb.ClusterStatus_value[j.Status]),
		})
	}
	return &pb.GetClusterListReply{
		Ok:       true,
		Clusters: clusters,
	}, nil
}

func (c *Client) DeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	credentials, _, err := c.reconcileCredentials(in.Name, in.Aws)
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	result, err := c.cmaAWSClient.DeleteCluster(cmaaws.DeleteClusterInput{
		Name: in.Name,
		Credentials: cmaaws.Credentials{
			Region:          credentials.Region,
			SecretKeyID:     credentials.SecretKeyID,
			SecretAccessKey: credentials.SecretAccessKey,
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
		logrus.Errorf("Could not set Cluster CR into kubernetes, this is bad")
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

func (c *Client) reconcileCredentials(clusterName string, providedCredentials *pb.AWSCredentials) (credentials awsk8sutil.Credentials, updateCache bool, err error) {
	logrus.Errorf("Reconciling credentials")
	if providedCredentials != nil &&
		providedCredentials.Region != "" &&
		providedCredentials.SecretKeyId != "" &&
		providedCredentials.SecretAccessKey != "" {
		logrus.Errorf("Using provided credentials")
		return awsk8sutil.Credentials{
			Region:          providedCredentials.Region,
			SecretKeyID:     providedCredentials.SecretKeyId,
			SecretAccessKey: providedCredentials.SecretAccessKey,
		}, true, nil
	}
	cacheResult, err := c.secretClient.GetCredentials(clusterName)
	if err != nil {
		return awsk8sutil.Credentials{}, false, err
	}
	logrus.Errorf("Using cached credentials")
	return cacheResult, false, nil
}

func (c *Client) updateCachedCredentials(clusterName string, credentials awsk8sutil.Credentials) (err error) {
	logrus.Errorf("Updating cached credentials")
	return c.secretClient.UpdateOrCreateCredentials(clusterName, credentials)
}

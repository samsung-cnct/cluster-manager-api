package cluster_manager_api

import (
	"fmt"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaws"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/aws"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func (s *Server) UpdateAWSCredentials(ctx context.Context, in *pb.UpdateAWSCredentialsMsg) (*pb.UpdateAWSCredentialsReply, error) {
	awsSecretClient, err := awsk8sutil.CreateFromDefaults()
	if err != nil {
		return &pb.UpdateAWSCredentialsReply{}, err
	}
	// TODO Add an UPDATE command
	awsSecretClient.DeleteCredentials(in.Name)
	err = awsSecretClient.CreateCredentials(in.Name, awsk8sutil.Credentials{
		Region:          in.Credentials.Region,
		SecretKeyID:     in.Credentials.SecretKeyId,
		SecretAccessKey: in.Credentials.SecretAccessKey,
	})
	if err != nil {
		return &pb.UpdateAWSCredentialsReply{}, err
	}
	return &pb.UpdateAWSCredentialsReply{Ok: true}, nil
}

func awsGetClient() (cmaaws.ClientInterface, error) {
	hostname := viper.GetString("cmaaws-endpoint")
	if hostname == "" {
		return nil, fmt.Errorf("aws support is not enabled")
	}
	insecure := viper.GetBool("cmaaws-insecure")
	return cmaaws.CreateNewClient(hostname, insecure)
}

func awsCreateCluster(in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	var instanceGroups []cmaaws.InstanceGroup
	client, err := awsGetClient()
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}
	defer client.Close()
	awsSecretClient, err := awsk8sutil.CreateFromDefaults()
	if err != nil {
		return &pb.CreateClusterReply{}, err
	}
	for _, j := range in.Provider.GetAws().InstanceGroups {
		instanceGroups = append(instanceGroups, cmaaws.InstanceGroup{
			Type:        j.Type,
			MinQuantity: int(j.MinQuantity),
			MaxQuantity: int(j.MaxQuantity),
		})
	}
	result, err := client.CreateCluster(cmaaws.CreateClusterInput{
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
	err = awsSecretClient.CreateCredentials(in.Name, awsk8sutil.Credentials{
		Region:          in.Provider.GetAws().Credentials.Region,
		SecretKeyID:     in.Provider.GetAws().Credentials.SecretKeyId,
		SecretAccessKey: in.Provider.GetAws().Credentials.SecretAccessKey,
	})

	if err != nil {
		// TODO Unsure what to do if we suddenly can't persist the credentails to kubernetes
		// TODO Going to log for now
		logger.Errorf("Could not set AWS credentials into kubernetes, this is bad")
	}

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     result.Cluster.ID,
			Name:   result.Cluster.Name,
			Status: result.Cluster.Status,
		},
	}, nil
}

func awsGetCluster(in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	client, err := awsGetClient()
	if err != nil {
		return &pb.GetClusterReply{}, err
	}
	defer client.Close()
	awsSecretClient, err := awsk8sutil.CreateFromDefaults()
	if err != nil {
		return &pb.GetClusterReply{}, err
	}
	credentials, err := awsSecretClient.GetCredentials(in.Name)
	if err != nil {
		return &pb.GetClusterReply{}, err
	}
	result, err := client.GetCluster(cmaaws.GetClusterInput{
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
	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         result.Cluster.ID,
			Name:       result.Cluster.Name,
			Status:     result.Cluster.Status,
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}, nil
}

func awsGetClusterList(in *pb.GetClusterListMsg) (*pb.GetClusterListReply, error) {
	var clusters []*pb.ClusterItem
	client, err := awsGetClient()
	if err != nil {
		return &pb.GetClusterListReply{}, err
	}
	defer client.Close()
	result, err := client.ListClusters(cmaaws.ListClusterInput{
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
			Status: j.Status,
		})
	}
	return &pb.GetClusterListReply{
		Ok:       true,
		Clusters: clusters,
	}, nil
}

func awsDeleteCluster(in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	client, err := awsGetClient()
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	defer client.Close()
	awsSecretClient, err := awsk8sutil.CreateFromDefaults()
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	credentials, err := awsSecretClient.GetCredentials(in.Name)
	if err != nil {
		return &pb.DeleteClusterReply{}, err
	}
	result, err := client.DeleteCluster(cmaaws.DeleteClusterInput{
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
	return &pb.DeleteClusterReply{
		Ok:     true,
		Status: result.Status,
	}, nil
}

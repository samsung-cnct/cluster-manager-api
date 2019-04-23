package cmaaws

import (
	"context"
	"crypto/tls"

	pb "github.com/samsung-cnct/cma-aws/pkg/generated/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.ClusterClient
}

func CreateNewClient(hostname string, insecure bool) (ClientInterface, error) {
	output := Client{}
	err := output.CreateNewClient(hostname, insecure)
	if err != nil {
		return nil, err
	}
	return &output, err
}

func (a *Client) CreateNewClient(hostname string, insecure bool) error {
	var err error
	if insecure {
		a.conn, err = grpc.Dial(hostname, grpc.WithInsecure())
		if err != nil {
			return err
		}
	} else {
		// If TLS is enabled, we're going to create credentials, also using built in certificates
		var tlsConf tls.Config
		creds := credentials.NewTLS(&tlsConf)

		a.conn, err = grpc.Dial(hostname, grpc.WithTransportCredentials(creds))
		if err != nil {
			return err
		}
	}
	a.client = pb.NewClusterClient(a.conn)
	return nil
}

func (a *Client) Close() error {
	return a.conn.Close()
}

func (a *Client) SetClient(client pb.ClusterClient) {
	a.client = client
}

func (a *Client) CreateCluster(input CreateClusterInput) (CreateClusterOutput, error) {
	var instanceGroups []*pb.CreateClusterAWSSpec_AWSInstanceGroup
	for _, j := range input.AWS.InstanceGroups {
		instanceGroups = append(instanceGroups, &pb.CreateClusterAWSSpec_AWSInstanceGroup{
			Type:        j.Type,
			MinQuantity: int32(j.MinQuantity),
			MaxQuantity: int32(j.MaxQuantity),
			// DesiredQuantity could be greater than MinQuantity with cluster wide autoscaling.
			DesiredQuantity: int32(j.MinQuantity),
			// Defaulting SshAccess to true until we have a field for that here.
			SshAccess: true,
		})
	}
	result, err := a.client.CreateCluster(context.Background(), &pb.CreateClusterMsg{
		Name: input.Name,
		Provider: &pb.CreateClusterProviderSpec{
			Name:       AWSProvider,
			K8SVersion: input.K8SVersion,
			Aws: &pb.CreateClusterAWSSpec{
				DataCenter: &pb.CreateClusterAWSSpec_AWSDataCenter{
					Region:            input.AWS.DataCenter.Region,
					AvailabilityZones: input.AWS.DataCenter.AvailabilityZones,
				},
				Credentials: &pb.AWSCredentials{
					Region:          input.AWS.Credentials.Region,
					SecretKeyId:     input.AWS.Credentials.SecretKeyID,
					SecretAccessKey: input.AWS.Credentials.SecretAccessKey,
				},
				Resources: &pb.CreateClusterAWSSpec_AWSPreconfiguredItems{
					VpcId:           input.AWS.PreconfiguredItems.VPCID,
					SecurityGroupId: input.AWS.PreconfiguredItems.SecurityGroupID,
					IamRoleArn:      input.AWS.PreconfiguredItems.IAMRoleARN,
				},
				InstanceGroups: instanceGroups,
			},
			HighAvailability: input.HighAvailability,
			NetworkFabric:    input.NetworkFabric,
		},
	})
	if err != nil {
		return CreateClusterOutput{}, err
	}
	output := CreateClusterOutput{
		Cluster: ClusterItem{
			ID:     result.Cluster.Id,
			Name:   result.Cluster.Name,
			Status: result.Cluster.Status.String(),
		},
	}
	return output, nil
}

func (a *Client) GetCluster(input GetClusterInput) (GetClusterOutput, error) {
	result, err := a.client.GetCluster(context.Background(), &pb.GetClusterMsg{
		Name: input.Name,
		Credentials: &pb.AWSCredentials{
			Region:          input.Credentials.Region,
			SecretKeyId:     input.Credentials.SecretKeyID,
			SecretAccessKey: input.Credentials.SecretAccessKey,
		},
		// Datacenter Region can be different from the credential Region, for now assuming the same.
		Region: input.Credentials.Region,
	})
	if err != nil {
		return GetClusterOutput{}, err
	}
	output := GetClusterOutput{
		Cluster: ClusterDetailItem{
			ID:         result.Cluster.Id,
			Name:       result.Cluster.Name,
			Status:     result.Cluster.Status.String(),
			Kubeconfig: result.Cluster.Kubeconfig,
		},
	}
	return output, nil
}

func (a *Client) DeleteCluster(input DeleteClusterInput) (DeleteClusterOutput, error) {
	result, err := a.client.DeleteCluster(context.Background(), &pb.DeleteClusterMsg{
		Name: input.Name,
		Credentials: &pb.AWSCredentials{
			Region:          input.Credentials.Region,
			SecretKeyId:     input.Credentials.SecretKeyID,
			SecretAccessKey: input.Credentials.SecretAccessKey,
		},
		// Datacenter Region can be different from the credential Region, for now assuming the same.
		Region: input.Credentials.Region,
	})
	if err != nil {
		return DeleteClusterOutput{}, err
	}
	output := DeleteClusterOutput{
		Status: result.Status,
	}
	return output, nil
}

func (a *Client) ListClusters(input ListClusterInput) (ListClusterOutput, error) {
	var clusters []ClusterItem
	result, err := a.client.GetClusterList(context.Background(), &pb.GetClusterListMsg{
		Credentials: &pb.AWSCredentials{
			Region:          input.Credentials.Region,
			SecretKeyId:     input.Credentials.SecretKeyID,
			SecretAccessKey: input.Credentials.SecretAccessKey,
		},
	})
	if err != nil {
		return ListClusterOutput{}, err
	}
	for _, j := range result.Clusters {
		clusters = append(clusters, ClusterItem{
			ID:     j.Id,
			Name:   j.Name,
			Status: j.Status.String(),
		})
	}
	output := ListClusterOutput{
		Clusters: clusters,
	}
	return output, nil
}

package aws

import (
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	BodyMissingError           = "No input data provided"
	CallbackMissingError       = "Callback is not provided"
	AWSArgumentsMissingError   = "AWS section is not provided"
	CredentialsMissingError    = "Credentials section is not provided"
	ResourcesMissingError      = "Resources section is not provided"
	DataCenterMissingError     = "DataCenter section is not provided"
	InstanceGroupsMissingError = "InstanceGroups section is empty"
)

func (c *Client) validateCreateClusterInput(in *pb.CreateClusterMsg) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, BodyMissingError)
	}
	if in.Callback == nil {
		return status.Error(codes.InvalidArgument, CallbackMissingError)
	}
	awsSection := in.Provider.GetAws()
	if awsSection == nil {
		return status.Error(codes.InvalidArgument, AWSArgumentsMissingError)
	}
	if awsSection.Credentials == nil {
		return status.Error(codes.InvalidArgument, CredentialsMissingError)
	}
	if awsSection.Resources == nil {
		return status.Error(codes.InvalidArgument, ResourcesMissingError)
	}
	if awsSection.DataCenter == nil {
		return status.Error(codes.InvalidArgument, DataCenterMissingError)
	}
	if len(awsSection.InstanceGroups) < 1 {
		return status.Error(codes.InvalidArgument, InstanceGroupsMissingError)
	}
	return nil
}

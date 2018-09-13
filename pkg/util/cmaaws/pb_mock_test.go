package cmaaws_test

import (
	"fmt"
	pb "gitlab.com/mvenezia/cma-aws/pkg/generated/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	MockBadCreateClusterErrorMessage  = "test create cluster error"
	MockBadGetClusterErrorMessage     = "test get cluster error"
	MockBadDeleteClusterErrorMessage  = "test delete cluster error"
	MockBadGetClusterListErrorMessage = "test get cluster list error"
	MockBadGetVersionErrorMessage     = "test get version error"
)

var (
	MockGoodCreateClusterReply = pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     "TestID",
			Name:   "TestName",
			Status: "TestStatus",
		},
	}
	MockGoodGetClusterReply = pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         "TestID",
			Name:       "TestName",
			Status:     "TestStatus",
			Kubeconfig: "TestKubeconfig",
		},
	}
	MockGoodDeleteClusterReply = pb.DeleteClusterReply{
		Ok:     true,
		Status: "TestDeleting",
	}
	MockGoodGetClusterListReply = pb.GetClusterListReply{
		Ok: true,
		Clusters: []*pb.ClusterItem{
			{Id: "Test1ID", Name: "Test1Name", Status: "Test1Status"},
			{Id: "Test2ID", Name: "Test2Name", Status: "Test2Status"},
		},
	}
	MockGoodGetVersionReply = pb.GetVersionReply{
		Ok:                 true,
		VersionInformation: &pb.GetVersionReply_VersionInformation{},
	}

	MockBadCreateClusterError  = fmt.Errorf(MockBadCreateClusterErrorMessage)
	MockBadGetClusterError     = fmt.Errorf(MockBadGetClusterErrorMessage)
	MockBadDeleteClusterError  = fmt.Errorf(MockBadDeleteClusterErrorMessage)
	MockBadGetClusterListError = fmt.Errorf(MockBadGetClusterListErrorMessage)
	MockBadGetVersionError     = fmt.Errorf(MockBadGetVersionErrorMessage)
)

type MockGoodCMAAWS struct{}

func (m *MockGoodCMAAWS) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg, opts ...grpc.CallOption) (*pb.CreateClusterReply, error) {
	return &MockGoodCreateClusterReply, nil
}

func (m *MockGoodCMAAWS) GetCluster(ctx context.Context, in *pb.GetClusterMsg, opts ...grpc.CallOption) (*pb.GetClusterReply, error) {
	return &MockGoodGetClusterReply, nil
}

func (m *MockGoodCMAAWS) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg, opts ...grpc.CallOption) (*pb.DeleteClusterReply, error) {
	return &MockGoodDeleteClusterReply, nil
}

func (m *MockGoodCMAAWS) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg, opts ...grpc.CallOption) (*pb.GetClusterListReply, error) {
	return &MockGoodGetClusterListReply, nil
}

func (m *MockGoodCMAAWS) GetVersionInformation(ctx context.Context, in *pb.GetVersionMsg, opts ...grpc.CallOption) (*pb.GetVersionReply, error) {
	return &MockGoodGetVersionReply, nil
}

type MockBadCMAAWS struct{}

func (m *MockBadCMAAWS) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg, opts ...grpc.CallOption) (*pb.CreateClusterReply, error) {
	return &pb.CreateClusterReply{}, MockBadCreateClusterError
}

func (m *MockBadCMAAWS) GetCluster(ctx context.Context, in *pb.GetClusterMsg, opts ...grpc.CallOption) (*pb.GetClusterReply, error) {
	return &pb.GetClusterReply{}, MockBadGetClusterError
}

func (m *MockBadCMAAWS) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg, opts ...grpc.CallOption) (*pb.DeleteClusterReply, error) {
	return &pb.DeleteClusterReply{}, MockBadDeleteClusterError
}

func (m *MockBadCMAAWS) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg, opts ...grpc.CallOption) (*pb.GetClusterListReply, error) {
	return &pb.GetClusterListReply{}, MockBadGetClusterListError
}

func (m *MockBadCMAAWS) GetVersionInformation(ctx context.Context, in *pb.GetVersionMsg, opts ...grpc.CallOption) (*pb.GetVersionReply, error) {
	return &pb.GetVersionReply{}, MockBadGetVersionError
}

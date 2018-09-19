package apiserver

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	service "github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/generated/api"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/website"
)

var (
	logger loggo.Logger
)

type ServerOptions struct {
	PortNumber int
}

func AddServersToMux(tcpMux cmux.CMux, options *ServerOptions) {
	logger = util.GetModuleLogger("pkg.apiserver", loggo.INFO)
	addGRPCServer(tcpMux)
	addRestAndWebsite(tcpMux, options.PortNumber)
}

func addGRPCServer(tcpMux cmux.CMux) {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterClusterServer(grpcServer, newgRPCServiceServer())
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	grpcListener := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	// Start servers
	go func() {
		logger.Infof("Starting gRPC Server")
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Criticalf("Unable to start external gRPC server")
		}
	}()
}

func addRestAndWebsite(tcpMux cmux.CMux, grpcPortNumber int) {
	httpListener := tcpMux.Match(cmux.HTTP1Fast())

	go func() {
		router := http.NewServeMux()
		website.AddWebsiteHandles(router)
		addgRPCRestGateway(router, grpcPortNumber)
		httpServer := http.Server{
			Handler: router,
		}
		logger.Infof("Starting HTTP/1 Server")
		httpServer.Serve(httpListener)
	}()

}

func addgRPCRestGateway(router *http.ServeMux, grpcPortNumber int) {
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()
	pb.RegisterClusterHandlerFromEndpoint(context.Background(), gwmux, "localhost:"+strconv.Itoa(grpcPortNumber), dopts)
	router.Handle("/api/", gwmux)
}

func newgRPCServiceServer() *service.Server {
	// TODO Handle Errors better here
	server, _ := service.NewServerFromDefaults()
	return server
}

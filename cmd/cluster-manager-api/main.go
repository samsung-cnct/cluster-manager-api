package main

import (
	"fmt"
	"log"
	"mime"
	"net"
	"net/http"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	service "github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api"
	"golang.org/x/net/context"

	"github.com/soheilhy/cmux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/swagger"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/swaggerjson"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/protobuf"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/homepage"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webserver on https://localhost:10000",
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func newServer() *service.Server {
	return new(service.Server)
}

func main() {
	port := 9050
	ctx := context.Background()

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	tcpMux := cmux.New(conn)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	dopts := []grpc.DialOption{grpc.WithInsecure()}

	pb.RegisterClusterServer(grpcServer, newServer())

	log.Println("Hi Mom")

	grpcListener := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpListener := tcpMux.Match(cmux.HTTP1Fast())
	// Start servers
	go func() {
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalln("Unable to start external gRPC server")
		}
	}()
	go func() {
		router := http.NewServeMux()
		//router.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		//	io.Copy(w, strings.NewReader(pb.APISwaggerJSON))
		//})
		//router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		//	http.Redirect(w, req, "index.html", 301)
		//})
		serveSwagger(router)
		serveSwaggerJSON(router)
		serveProtoBuf(router)
		serveHomepage(router)
		gwmux := runtime.NewServeMux()
		pb.RegisterClusterHandlerFromEndpoint(ctx, gwmux, "localhost:9050", dopts)
		router.Handle("/api/", gwmux)
		httpServer := http.Server{
			Handler: router,
		}
		httpServer.Serve(httpListener)
	}()

	tcpMux.Serve()
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveSwaggerJSON(mux *http.ServeMux) {
	mime.AddExtensionType(".json", "application/json")

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swaggerjson.Asset,
		AssetDir: swaggerjson.AssetDir,
	})
	prefix := "/swagger/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveProtoBuf(mux *http.ServeMux) {

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    protobuf.Asset,
		AssetDir: protobuf.AssetDir,
	})
	prefix := "/protobuf/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveHomepage(mux *http.ServeMux) {

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    homepage.Asset,
		AssetDir: homepage.AssetDir,
	})
	prefix := "/"
	mux.Handle(prefix, http.StripPrefix("/", fileServer))
}


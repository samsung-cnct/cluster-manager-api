package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"mime"
	"net"
	"net/http"

	service "github.com/samsung-cnct/cluster-manager-api/internal/cluster-manager-api"
	pb "github.com/samsung-cnct/cluster-manager-api/pkg/api"
	"golang.org/x/net/context"

	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/homepage"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/protobuf"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/swagger"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/swaggerjson"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/venezia/redis-operator/pkg/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/juju/loggo"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/samsung-cnct/cluster-controller/pkg/client/clientset/versioned"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

	clusterController "github.com/samsung-cnct/cluster-manager-api/pkg/cluster-controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

var (
	logger loggo.Logger
	config *rest.Config
)

const (
	kubeconfigDir  = ".kube"
	kubeconfigFile = "config"
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
	var err error
	logger := util.GetModuleLogger("cmd.redis-operator", loggo.INFO)
	viperInit()

	// get flags
	portNumber := viper.GetInt("port")
	kubeconfigLocation := viper.GetString("kubeconfig")

	// Debug for now
	logger.Infof("Parsed Variables: \n  Port: %d \n  Kubeconfig: %s", portNumber, kubeconfigLocation)

	if kubeconfigLocation != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigLocation)
		if err != nil {
			logErrorAndExit(err)
		}
	} else {
		configPath := filepath.Join(homeDir(), kubeconfigDir, kubeconfigFile)
		if _, err := os.Stat(configPath); err == nil {
			config, err = clientcmd.BuildConfigFromFlags("", configPath)
		} else {
			config, err = rest.InClusterConfig()
		}
	}

	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logErrorAndExit(err)
	}

	clusterControllerClient := clusterController.New(clusterController.Config{
		KubeCli:     clientSet,
		KubeExtCli:  apiextensionsclient.NewForConfigOrDie(config),
		KrakenCRCli: versioned.NewForConfigOrDie(config),
	})

	pods, _ := clusterControllerClient.KubeCli.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		logErrorAndExit(err)
	}

	logger.Infof("There are %d pods in the cluster", len(pods.Items))

	ctx := context.Background()

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
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
		pb.RegisterClusterHandlerFromEndpoint(ctx, gwmux, "localhost:"+strconv.Itoa(portNumber), dopts)
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

func viperInit() {
	viper.SetEnvPrefix("clustermanagerapi")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// using standard library "flag" package
	flag.Int("port", 9050, "Port to listen on")
	flag.String("kubeconfig", "", "Location of kubeconfig file")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.AutomaticEnv()
}

func logErrorAndExit(err error) {
	logger.Criticalf("error: %s", err)
	os.Exit(1)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

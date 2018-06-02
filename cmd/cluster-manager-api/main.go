package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/samsung-cnct/cma-operator/pkg/util"
	"github.com/samsung-cnct/cma-operator/pkg/util/k8sutil"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/apiserver"
	"k8s.io/client-go/rest"
	"github.com/spf13/cobra"
	"encoding/json"
	"github.com/samsung-cnct/cluster-manager-api/pkg/version"
)

var (
	logger loggo.Logger
	config *rest.Config
	OutputFormat string
)


func main() {
	var rootCmd = &cobra.Command{
		Use:   "cluster-manager-api",
		Short: "The cluster manager API",
		Long: `The Cluster Manager API`,
		Run: func(cmd *cobra.Command, args []string) {
			doSomething()
		},
	}

	cobra.OnInitialize()
	viperInit()
	rootCmd.AddCommand(generateVersionCmd())
	rootCmd.Execute()
}


func doSomething() {
	var err error
	logger := util.GetModuleLogger("cmd.cma-operator", loggo.INFO)

	// get flags
	portNumber := viper.GetInt("port")
	kubeconfigLocation := viper.GetString("kubeconfig")

	// Debug for now
	logger.Infof("Parsed Variables: \n  Port: %d \n  Kubeconfig: %s", portNumber, kubeconfigLocation)

	k8sutil.KubeConfigLocation = kubeconfigLocation
	k8sutil.DefaultConfig, err = k8sutil.GenerateKubernetesConfig()

	if err != nil {
		logger.Infof("Was unable to generate a valid kubernetes default config, some functionality may be broken.  Error was %v", err)
	}

	var wg sync.WaitGroup
	stop := make(chan struct{})

	logger.Infof("Creating Web Server")
	tcpMux := createWebServer(&apiserver.ServerOptions{PortNumber: portNumber})
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Infof("Starting to serve requests on port %d", portNumber)
		tcpMux.Serve()
	}()

	<-stop
	logger.Infof("Wating for controllers to shut down gracefully")
	wg.Wait()
}

func createWebServer(options *apiserver.ServerOptions) cmux.CMux {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", options.PortNumber))
	if err != nil {
		panic(err)
	}
	tcpMux := cmux.New(conn)

	apiserver.AddServersToMux(tcpMux, options)

	return tcpMux
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


func generateVersionCmd() *cobra.Command{
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Returns version information",
		Long: `Find out the version, git commit, etc of the build`,
		Run: func(cmd *cobra.Command, args []string) {
			info := version.Get()
			if OutputFormat == "json" {
				versionString, _ := json.Marshal(info)
				fmt.Printf("%s\n", versionString)
			} else {
				fmt.Printf("Version Information:\n")
				fmt.Printf("\tGit Data:\n")
				fmt.Printf("\t\tTagged Version:\t%s\n", info.GitVersion)
				fmt.Printf("\t\tHash:\t\t%s\n", info.GitCommit)
				fmt.Printf("\t\tTree State:\t%s\n", info.GitTreeState)
				fmt.Printf("\t\tBuild Date:\t%s\n", info.BuildDate)
				fmt.Printf("\tBuild Data:\n")
				fmt.Printf("\t\tBuild Date:\t%s\n", info.BuildDate)
				fmt.Printf("\t\tGo Version:\t%s\n", info.GoVersion)
				fmt.Printf("\t\tCompiler:\t%s\n", info.Compiler)
				fmt.Printf("\t\tPlatform:\t%s\n\n", info.Platform)
			}
		},
	}

	versionCmd.Flags().StringVarP(&OutputFormat, "output", "o", "text", "json or text")
	return versionCmd
}

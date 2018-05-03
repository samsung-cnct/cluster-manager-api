package k8sutil

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
	"os"
)

var (
	KubeConfigLocation string
	DefaultConfig *rest.Config
)

const (
	kubeconfigDir  = ".kube"
	kubeconfigFile = "config"
)

func GenerateKubernetesConfig() (*rest.Config, error) {
	var config *rest.Config
	var err error

	if KubeConfigLocation != "" {
		config, err = clientcmd.BuildConfigFromFlags("", KubeConfigLocation)
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
	return config, err
}
package cmak8sutil

import (
	"fmt"
	cmaVersioned "github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/clientset/versioned"
	"github.com/samsung-cnct/cma-operator/pkg/generated/cma/client/clientset/versioned/typed/cma/v1alpha1"
	"github.com/spf13/viper"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

const (
	kubeconfigDir  = ".kube"
	kubeconfigFile = "config"
)

func CreateFromDefaults() (ClientInterface, error) {
	output := Client{}
	err := output.CreateNewClients()
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (c *Client) CreateNewClients() error {
	var err error
	if c.config == nil {
		err = c.createNewConfig()
		if err != nil {
			return err
		}
	}
	cmaClient, err := cmaVersioned.NewForConfig(c.config)
	if err != nil {
		return err
	}
	c.applicationClient = cmaClient.CmaV1alpha1().SDSApplications(viper.GetString("kubernetes-namespace"))
	c.clusterClient = cmaClient.CmaV1alpha1().SDSClusters(viper.GetString("kubernetes-namespace"))
	c.packageManagerClient = cmaClient.CmaV1alpha1().SDSPackageManagers(viper.GetString("kubernetes-namespace"))
	return err
}

func (c *Client) SetApplicationClient(client v1alpha1.SDSApplicationInterface) {
	c.applicationClient = client
}

func (c *Client) SetClusterClient(client v1alpha1.SDSClusterInterface) {
	c.clusterClient = client
}

func (c *Client) SetPackageManagerClient(client v1alpha1.SDSPackageManagerInterface) {
	c.packageManagerClient = client
}

func (c *Client) createNewConfig() error {
	var err error
	if c.kubeConfigLocation != "" {
		c.config, err = clientcmd.BuildConfigFromFlags("", c.kubeConfigLocation)
	} else {
		configPath := filepath.Join(homeDir(), kubeconfigDir, kubeconfigFile)
		_, err := os.Stat(configPath)
		if err == nil {
			c.config, err = clientcmd.BuildConfigFromFlags("", configPath)
		} else {
			c.config, err = rest.InClusterConfig()
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func (c *Client) SetConfig(config *rest.Config) {
	c.config = config
}

func (c *Client) getAdjustedName(name string, clusterName string) string {
	return name + "-" + clusterName
}

func (c *Client) removeAdjustedName(name string, clusterName string) (string, error) {
	inputLength := len(name)
	suffixLength := len(clusterName)
	if suffixLength >= inputLength {
		return "", fmt.Errorf("input string is longer than suffix")
	}
	if name[(inputLength-suffixLength):] != clusterName {
		return "", fmt.Errorf("input string does not have the suffix")
	}
	return name[0:(inputLength - suffixLength - 1)], nil
}

func (c *Client) getAdjustedApplicationName(name string, packageManagerName string, clusterName string) string {
	return name + "-" + packageManagerName + "-" + clusterName
}

func (c *Client) removeAdjustedApplicationName(name string, packageManagerName string, clusterName string) (string, error) {
	inputLength := len(name)
	suffixLength := len(packageManagerName + "-" + clusterName)
	if suffixLength >= inputLength {
		return "", fmt.Errorf("input string is longer than suffix")
	}
	if name[(inputLength-suffixLength):] != packageManagerName+"-"+clusterName {
		return "", fmt.Errorf("input string does not have the suffix")
	}
	return name[0:(inputLength - suffixLength - 1)], nil
}

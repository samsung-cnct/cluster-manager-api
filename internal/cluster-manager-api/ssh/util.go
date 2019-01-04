package ssh

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmassh"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
	"github.com/spf13/viper"
)

const (
	CMASshEndpointViperVariableName = "cmassh-endpoint"
	CMASshInsecureViperVariableName = "cmassh-insecure"
	NotEnabledErrorMessage             = "ssh support is not enabled"
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

	// Adding the CMAAKS Client
	c.cmaSshClient, err = getCMASshClient()
	if err != nil {
		return err
	}

	c.cmaK8sClient, err = cmak8sutil.CreateFromDefaults()
	if err != nil {
		// Closing because we created the client but then errored
		c.Close()
		return err
	}

	return nil
}

func getCMASshClient() (cmassh.ClientInterface, error) {
	if IsEnabled() != true {
		return nil, fmt.Errorf(NotEnabledErrorMessage)
	}

	hostname := viper.GetString(CMASshEndpointViperVariableName)
	insecure := viper.GetBool(CMASshInsecureViperVariableName)
	return cmassh.CreateNewClient(hostname, insecure)
}

func (c *Client) SetCMASshClient(client cmassh.ClientInterface) {
	c.cmaSshClient = client
}

func (c *Client) SetCMAK8sClient(client cmak8sutil.ClientInterface) {
	c.cmaK8sClient = client
}

func (c *Client) Close() error {
	return c.cmaSshClient.Close()
}

func IsEnabled() bool {
	if viper.GetString(CMASshEndpointViperVariableName) == "" {
		return false
	}
	return true
}

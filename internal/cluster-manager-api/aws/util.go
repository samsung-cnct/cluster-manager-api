package aws

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaws"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/aws"
	"github.com/spf13/viper"
)

const (
	CMAAWSEndpointViperVariableName = "cmaaws-endpoint"
	CMAAWSInsecureViperVariableName = "cmaaws-insecure"
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

	// Adding the CMAAWS Client
	c.cmaAWSClient, err = getCMAAWSClient()
	if err != nil {
		return err
	}

	c.secretClient, err = awsk8sutil.CreateFromDefaults()
	if err != nil {
		// Closing because we created the client but then errored
		c.Close()
		return err
	}

	return nil
}

func getCMAAWSClient() (cmaaws.ClientInterface, error) {
	hostname := viper.GetString(CMAAWSEndpointViperVariableName)
	if hostname == "" {
		return nil, fmt.Errorf("azure support is not enabled")
	}
	insecure := viper.GetBool(CMAAWSInsecureViperVariableName)
	return cmaaws.CreateNewClient(hostname, insecure)
}

func (c *Client) Close() error {
	return c.cmaAWSClient.Close()
}

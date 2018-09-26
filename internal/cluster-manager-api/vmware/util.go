package vmware

import (
	"fmt"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmavmware"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil/cma"
	"github.com/spf13/viper"
)

const (
	CMAVMWareEndpointViperVariableName = "cmavmware-endpoint"
	CMAVMWareInsecureViperVariableName = "cmavmware-insecure"
	NotEnabledErrorMessage             = "vmware support is not enabled"
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
	c.cmaVMWareClient, err = getCMAVMWareClient()
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

func getCMAVMWareClient() (cmavmware.ClientInterface, error) {
	if IsEnabled() != true {
		return nil, fmt.Errorf(NotEnabledErrorMessage)
	}

	hostname := viper.GetString(CMAVMWareEndpointViperVariableName)
	insecure := viper.GetBool(CMAVMWareInsecureViperVariableName)
	return cmavmware.CreateNewClient(hostname, insecure)
}

func (c *Client) SetCMAVMWareClient(client cmavmware.ClientInterface) {
	c.cmaVMWareClient = client
}

func (c *Client) SetCMAK8sClient(client cmak8sutil.ClientInterface) {
	c.cmaK8sClient = client
}

func (c *Client) Close() error {
	return c.cmaVMWareClient.Close()
}

func IsEnabled() bool {
	if viper.GetString(CMAVMWareEndpointViperVariableName) == "" {
		return false
	}
	return true
}

package awsk8sutil

import (
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type Client struct {
	kubeConfigLocation string
	config             *rest.Config
	secretClient       v1.SecretInterface
}

type ClientInterface interface {
	CreateCredentials(name string, credentials Credentials) error
	GetCredentials(name string) (Credentials, error)
	DeleteCredentials(name string) error
	ListCredentials() ([]string, error)

	CreateNewClients() error
	SetConfig(config *rest.Config)
	SetSecretClient(client v1.SecretInterface)
}

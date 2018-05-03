package k8sutil

import (
	"github.com/juju/loggo"
	"os"
)

var (
	logger loggo.Logger
)

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

package k8sutil

import (
	"github.com/juju/loggo"
	"os"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
)

var (
	logger loggo.Logger
)

func SetLogger() {
	logger = util.GetModuleLogger("pkg.util.k8sutil", loggo.INFO)
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

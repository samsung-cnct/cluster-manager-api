package cmaaws_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCmaaws(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cmaaws Suite")
}

package twiml_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwiml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Twiml Suite")
}

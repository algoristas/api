package results_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestResults(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Results Suite")
}

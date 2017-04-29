package problems_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProblems(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Problems Suite")
}

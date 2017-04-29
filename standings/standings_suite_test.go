package standings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStandings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Standings Suite")
}

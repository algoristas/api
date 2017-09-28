package standings_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/algoristas/api/router"
)

type StubbedDataProvider struct{}

func (t *StubbedDataProvider) GetStandings() ([]byte, error) {
	return []byte(`{"users":[]}`), nil
}

var _ = Describe("Standings", func() {
	var ts = httptest.NewServer(router.Wire(router.Dependencies{
		StandingsDataProvider: &StubbedDataProvider{},
	}))

	It("should return standings data", func() {
		resp, err := http.Get(ts.URL + "/v1/standings")
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())
		Expect(data).To(ContainSubstring(`{"users":[]}`))
	})
})

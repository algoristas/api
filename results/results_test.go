package results_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/algoristas/api/router"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type StubbedDataProvider struct {
}

func (t *StubbedDataProvider) GetResults() ([]byte, error) {
	return []byte(`{"users":[]}`), nil
}

var _ = Describe("Results", func() {
	var ts = httptest.NewServer(router.Wire(router.Dependencies{
		ResultsDataProvider: &StubbedDataProvider{},
	}))

	It("should return results", func() {
		resp, err := http.Get(ts.URL + "/v1/results")
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())
		Expect(data).To(ContainSubstring(`{"users":[]}`))
	})
})

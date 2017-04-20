package problems_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/algoristas/api/router"
)

type DAOStub struct {
}

func (t *DAOStub) GetSets() ([]byte, error) {
	return []byte(`{"weeks":[]}`), nil
}

var _ = Describe("Problems", func() {
	var ts = httptest.NewServer(router.Wire(router.Dependencies{
		ProblemsDAO: &DAOStub{},
	}))

	It("returns problem list", func() {
		log.Printf("URL: %s", ts.URL)
		resp, err := http.Get(ts.URL + "/v1/problems/sets")
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())
		Expect(data).To(ContainSubstring(`{"weeks":[]}`))
	})
})

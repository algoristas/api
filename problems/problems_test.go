package problems_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/algoristas/api/problems"
	"github.com/algoristas/api/router"
	"github.com/algoristas/api/test"
	"github.com/algoristas/api/users"
)

// See fixtures/users.json
const codeforcesID = 3

type ProblemResponse struct {
	StatusCode uint             `json:"statusCode"`
	Errors     []string         `json:"errors"`
	Data       problems.Problem `json:"data"`
}

type StubbedDataProvider struct {
}

func (t *StubbedDataProvider) GetSets() ([]byte, error) {
	return []byte(`{"weeks":[]}`), nil
}

func (t *StubbedDataProvider) FindProblem(userID string, problemID uint) (*problems.Problem, error) {
	return nil, errors.New("something went wrong")
}

var _ = Describe("Problems", func() {
	var ts = httptest.NewServer(router.Wire(router.Dependencies{
		ProblemsDataProvider: &StubbedDataProvider{},
		UsersDataProvider:    users.NewDataProvider(),
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

	Describe("Codeforces", func() {
		It("should return error if user is not found", func() {
			resp, err := http.Get(ts.URL + "/v1/users/nobody/problems/1")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

			errorResponse := test.ReadError(resp)
			Expect(errorResponse.StatusCode).To(Equal(http.StatusNotFound))
			Expect(errorResponse.Message).To(Equal("User not found"))
		})

		It("should return error if the problem can not be retrieved", func() {
			resp, err := http.Get(ts.URL + "/v1/users/rendon/problems/1")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))

			errorResponse := test.ReadError(resp)
			Expect(errorResponse.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(errorResponse.Message).To(Equal("Failed to retrieve problem"))
		})

		// This test is more of an integration test, keeping it here for now
		It("should return problem data", func() {
			var ts = httptest.NewServer(router.Wire(router.Dependencies{
				ProblemsDataProvider: problems.NewDataProvider(),
				UsersDataProvider:    users.NewDataProvider(),
			}))

			resp, err := http.Get(ts.URL + "/v1/users/rendon/problems/1")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			var problemResponse ProblemResponse
			Expect(decoder.Decode(&problemResponse)).To(BeNil())

			problem := problemResponse.Data
			Expect(problem.ID).To(Equal(uint(1)))
			Expect(problem.OwnerID).To(Equal(codeforcesID))
			Expect(problem.Title).To(Equal("Kefa and Park"))
			Expect(problem.Source).To(Equal("http://codeforces.com/problemset/problem/580/C"))
			Expect(problem.HasSolved).To(BeTrue())
			Expect(problem.HasTried).To(BeTrue())
		})
	})
})

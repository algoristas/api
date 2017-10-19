package users_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/algoristas/api/model"
	"github.com/algoristas/api/router"
	"github.com/algoristas/api/users"
)

var usersDataProvider = users.NewDataProvider()

// StubbedDataProvider implements users.DataProvider to test uncommon scenarios.
type StubbedDataProvider struct{}

func (t *StubbedDataProvider) FindUser(userName string) (*users.User, error) {
	return nil, errors.New("Something went wrong!")
}

func (t *StubbedDataProvider) FindUserByID(id uint) (*users.User, error) {
	return nil, errors.New("Something went wrong!")
}

func (t *StubbedDataProvider) FindUserByUserName(userName string) (*users.User, error) {
	return nil, errors.New("Something went wrong!")
}

var _ = Describe("Users", func() {
	Describe("Find user by ID", func() {
		It("should return user if it exists", func() {
			user, err := usersDataProvider.FindUserByID(1)
			Expect(err).To(BeNil())
			Expect(user.ID).To(Equal(uint(1)))
			Expect(user.UserName).To(Equal("alice"))
		})

		It("should return error if user does not exist", func() {
			_, err := usersDataProvider.FindUserByID(11)
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("Find user by user name", func() {
		It("should return user if it exists", func() {
			user, err := usersDataProvider.FindUserByUserName("bob")
			Expect(err).To(BeNil())
			Expect(user.ID).To(Equal(uint(2)))
			Expect(user.UserName).To(Equal("bob"))
		})

		It("should return error if user does not exist", func() {
			_, err := usersDataProvider.FindUserByUserName("alice_and_bob")
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("Users controller", func() {
		var ts = httptest.NewServer(router.Wire(router.Dependencies{
			UsersDataProvider: users.NewDataProvider(),
		}))

		It("should return error if user does not exist", func() {
			resp, err := http.Get(ts.URL + "/v1/users/nouser")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

			defer resp.Body.Close()
			var decoder = json.NewDecoder(resp.Body)

			var errorResponse model.ErrorResponse
			Expect(decoder.Decode(&errorResponse)).To(BeNil())
			Expect(errorResponse.StatusCode).To(Equal(http.StatusNotFound))
			Expect(errorResponse.Message).To(Equal("User not found"))
		})

		It("should return user by ID", func() {
			resp, err := http.Get(ts.URL + "/v1/users/1")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			defer resp.Body.Close()
			var decoder = json.NewDecoder(resp.Body)

			var user users.User
			Expect(decoder.Decode(&user)).To(BeNil())
			Expect(user.ID).To(Equal(uint(1)))
			Expect(user.UserName).To(Equal("alice"))
		})

		It("should return user by user name", func() {
			resp, err := http.Get(ts.URL + "/v1/users/bob")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			defer resp.Body.Close()
			var decoder = json.NewDecoder(resp.Body)

			var user users.User
			Expect(decoder.Decode(&user)).To(BeNil())
			Expect(user.ID).To(Equal(uint(2)))
			Expect(user.UserName).To(Equal("bob"))
		})

		It("should return internal server error if something goes wrong", func() {
			var ts = httptest.NewServer(router.Wire(router.Dependencies{
				UsersDataProvider: &StubbedDataProvider{},
			}))

			resp, err := http.Get(ts.URL + "/v1/users/bob")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))

			defer resp.Body.Close()
			var decoder = json.NewDecoder(resp.Body)

			var errorResponse model.ErrorResponse
			Expect(decoder.Decode(&errorResponse)).To(BeNil())
			Expect(errorResponse.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(errorResponse.Message).To(Equal("Internal Server Error"))
		})
	})
})

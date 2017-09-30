package users

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var ErrNotFound = errors.New("User not found")

// User describes a user.
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
}

// DefaultDataProvider implements the DataProvider interface.
type DefaultDataProvider struct{}

// FindUserByID finds user by its ID, returns error if the user does not exist
// or something goes wrong retrieving the user.
func (t *DefaultDataProvider) FindUserByID(id uint) (*User, error) {
	users, err := retrieveUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, ErrNotFound
}

// FindUserByUserName finds user by its user name, return error if the user does
// not exist or something goes wrong retrieving the user.
func (t *DefaultDataProvider) FindUserByUserName(userName string) (*User, error) {
	users, err := retrieveUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.UserName == userName {
			return &user, nil
		}
	}
	return nil, ErrNotFound
}

// retrieveUsers retrieves users from our database (a file for now).
func retrieveUsers() ([]User, error) {
	fileName := os.Getenv("APP_ROOT") + "/users/fixtures/users.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	return users, err
}

// NewDataProvider returns a new instance of DataProvider.
func NewDataProvider() DataProvider {
	return &DefaultDataProvider{}
}

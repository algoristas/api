package problems

import (
	"io/ioutil"
	"os"
)

// DefaultDataProvider implements the DataProvider interface.
type DefaultDataProvider struct {
}

// GetSets returns all the problems in the database.
func (t *DefaultDataProvider) GetSets() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/problems/fixtures/problemsets.json"
	return ioutil.ReadFile(fileName)
}

// NewDataProvider returns a new DataProvider instance.
func NewDataProvider() DataProvider {
	return &DefaultDataProvider{}
}

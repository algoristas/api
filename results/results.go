package results

import (
	"io/ioutil"
	"os"
)

// DefaultDataProvider implements the DataProvider interface.
type DefaultDataProvider struct{}

// GetResults retrieves all contests results.
func (t *DefaultDataProvider) GetResults() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/results/fixtures/results.json"
	return ioutil.ReadFile(fileName)
}

// NewDataProvider returns a new DataProvider instance.
func NewDataProvider() DataProvider {
	return &DefaultDataProvider{}
}

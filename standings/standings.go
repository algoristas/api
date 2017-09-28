package standings

import (
	"io/ioutil"
	"os"
)

// DefaultDataProvider implements DataProvider interface.
type DefaultDataProvider struct{}

// GetStandings retrieves the standings for all contestants up to this moment.
func (t *DefaultDataProvider) GetStandings() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/standings/fixtures/standings.json"
	return ioutil.ReadFile(fileName)
}

// NewDataProvider returns a new DataProvider instance.
func NewDataProvider() DataProvider {
	return &DefaultDataProvider{}
}

package results

import (
	"io/ioutil"
	"os"
)

// ResultsDAO implements the DAO interface.
type ResultsDAO struct{}

// GetResults retrieves all contests results.
func (t *ResultsDAO) GetResults() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/results/fixtures/results.json"
	return ioutil.ReadFile(fileName)
}

// NewResultsDAO returns a new object that implements the DAO interface.
func NewResultsDAO() DAO {
	return &ResultsDAO{}
}

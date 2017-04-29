package problems

import (
	"io/ioutil"
	"os"
)

// ProblemsDAO implements the DAO interface.
type ProblemsDAO struct {
}

// GetSets returns all the problems in the database.
func (t *ProblemsDAO) GetSets() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/problems/fixtures/problemsets.json"
	return ioutil.ReadFile(fileName)
}

// NewProblemsDAO returns a new object that implemts the DAO interface.
func NewProblemsDAO() DAO {
	return &ProblemsDAO{}
}

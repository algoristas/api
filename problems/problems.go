package problems

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// DefaultDataProvider implements the DataProvider interface.
type DefaultDataProvider struct {
}

// GetSets returns all the problems in the database.
func (t *DefaultDataProvider) GetSets() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/problems/fixtures/problemsets.json"
	return ioutil.ReadFile(fileName)
}

// FindProblem retrieves problem details for a given user and problem ID. For now it only works with
// Codeforces.
func (t *DefaultDataProvider) FindProblem(userID, problemID string) (*Problem, error) {

	// NOTE: Assumming this is a Codeforces ID
	contestID, err := strconv.Atoi(problemID[0 : len(problemID)-1])
	//index := problemID[len(problemID)-1:]

	url := fmt.Sprintf(
		"http://codeforces.com/api/contest.status?contestId=%d&handle=%s&showUnofficial=true",
		contestID, userID,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Data:\n%s\n", data)
	return &Problem{}, nil
}

// NewDataProvider returns a new DataProvider instance.
func NewDataProvider() DataProvider {
	return &DefaultDataProvider{}
}

// Problem describes a problem in our system.
type Problem struct {
	ID        string
	OwnerID   int
	Title     string
	Source    string
	HasSolved bool
	HasTried  bool
}

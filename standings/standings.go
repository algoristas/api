package standings

import (
	"io/ioutil"
	"os"
)

// StandingDAO implements DAO interface.
type StandingsDAO struct{}

// GetStandings retrieves the standings for all contestants up to this moment.
func (t *StandingsDAO) GetStandings() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/standings/fixtures/standings.json"
	return ioutil.ReadFile(fileName)
}

// NewStandingsDAO returns a new object that implements the DAO interface.
func NewStandingsDAO() DAO {
	return &StandingsDAO{}
}

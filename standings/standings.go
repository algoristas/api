package standings

import (
	"io/ioutil"
	"os"
)

func GetStandings() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/tests/fixtures/standings.json"
	return ioutil.ReadFile(fileName)
}

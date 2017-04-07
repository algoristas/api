package standings

import (
	"io/ioutil"
	"log"
	"os"
)

func GetStandings() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/tests/fixtures/standings.json"
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Failed to read file: %s", err)
		return nil, err
	}
	return buf, nil
}

package results

import (
	"io/ioutil"
	"os"
)

func GetResults() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/tests/fixtures/results.json"
	return ioutil.ReadFile(fileName)
}

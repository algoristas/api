package problems

import (
	"io/ioutil"
	"os"
)

func GetSets() ([]byte, error) {
	fileName := os.Getenv("APP_ROOT") + "/tests/fixtures/problemsets.json"
	return ioutil.ReadFile(fileName)
}

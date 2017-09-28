package problems

import (
	"fmt"
	"log"
	"net/http"
)

// Controller describes controller for requests at /problems/.
type Controller struct {
	dataProvider DataProvider
}

// SetIndex handles /problems/sets endpoint, returns all problems.
func (t *Controller) SetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := t.dataProvider.GetSets()
	if err != nil {
		log.Printf("Failed to retrieve results: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// NewController returns a new initialized Controller.
func NewController(datProvider DataProvider) *Controller {
	return &Controller{
		dataProvider: datProvider,
	}
}

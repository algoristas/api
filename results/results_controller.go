package results

import (
	"fmt"
	"log"
	"net/http"
)

// ResultsController describes controller for requests at /results/.
type ResultsController struct {
	dao DAO
}

// Index handles /results/ endpoint, returns all results.
func (t *ResultsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := t.dao.GetResults()
	if err != nil {
		log.Printf("Failed to retrieve results: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// NewController returns an initialized ResultsController.
func NewController(resultsDAO DAO) *ResultsController {
	return &ResultsController{
		dao: resultsDAO,
	}
}

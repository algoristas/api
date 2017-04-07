package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/algoristas/api/results"
)

type ResultsController struct {
}

func (t *ResultsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := results.GetResults()
	if err != nil {
		log.Printf("Failed to retrieve results: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func NewResultsController() *ResultsController {
	return &ResultsController{}
}

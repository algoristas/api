package controllers

import (
	"fmt"
	"net/http"
)

type ResultsController struct {
}

func (t *ResultsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{}`)
}

func NewResultsController() *ResultsController {
	return &ResultsController{}
}

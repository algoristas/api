package controllers

import (
	"fmt"
	"net/http"
)

type StandingsController struct {
}

func (t *StandingsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{}`)
}

func NewStandingsController() *StandingsController {
	return &StandingsController{}
}

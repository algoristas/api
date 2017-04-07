package controllers

import (
	"fmt"
	"net/http"

	"github.com/algoristas/api/standings"
)

type StandingsController struct {
}

func (t *StandingsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := standings.GetStandings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func NewStandingsController() *StandingsController {
	return &StandingsController{}
}

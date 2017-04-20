package standings

import (
	"fmt"
	"net/http"
)

// StandingsController describes controller for requests at /standings/.
type StandingsController struct {
	dao DAO
}

// Index handles /standings endpoint, returns the list of standings.
func (t *StandingsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := t.dao.GetStandings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// NewController returns an initialized StandingsController.
func NewController(standingsDAO DAO) *StandingsController {
	return &StandingsController{
		dao: standingsDAO,
	}
}

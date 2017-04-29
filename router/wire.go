package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/algoristas/api/problems"
	"github.com/algoristas/api/results"
	"github.com/algoristas/api/standings"
)

// Dependencies contains every dependency used by the application.
type Dependencies struct {
	StandingsDAO standings.DAO
	ProblemsDAO  problems.DAO
	ResultsDAO   results.DAO
}

// Wire returns an http.Handler with all the API endpoints configured using the provided
// dependencies.
func Wire(deps Dependencies) http.Handler {
	r := mux.NewRouter()

	standingsController := standings.NewController(deps.StandingsDAO)
	r.HandleFunc("/v1/standings", standingsController.Index)

	resultsController := results.NewController(deps.ResultsDAO)
	r.HandleFunc("/v1/results", resultsController.Index)

	problemsController := problems.NewController(deps.ProblemsDAO)
	r.HandleFunc("/v1/problems/sets", problemsController.SetIndex)
	return r
}

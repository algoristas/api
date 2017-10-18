package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/algoristas/api/problems"
	"github.com/algoristas/api/results"
	"github.com/algoristas/api/standings"
	"github.com/algoristas/api/users"
)

// Dependencies contains every dependency used by the application.
type Dependencies struct {
	StandingsDataProvider standings.DataProvider
	ProblemsDataProvider  problems.DataProvider
	ResultsDataProvider   results.DataProvider
	UsersDataProvider     users.DataProvider
}

// Wire returns an http.Handler with all the API endpoints configured using the provided
// dependencies.
func Wire(deps Dependencies) http.Handler {
	r := httprouter.New()

	standingsController := standings.NewController(deps.StandingsDataProvider)
	r.GET("/v1/standings", standingsController.Index)

	resultsController := results.NewController(deps.ResultsDataProvider)
	r.GET("/v1/results", resultsController.Index)

	problemsController := problems.NewController(deps.ProblemsDataProvider, deps.UsersDataProvider)
	r.GET("/v1/problems/sets", problemsController.SetIndex)
	r.GET("/v1/users/:userID/problems/:problemID", problemsController.GetProblem)

	usersController := users.NewController(deps.UsersDataProvider)
	r.GET("/v1/users/:userID", usersController.GetUser)
	return r
}

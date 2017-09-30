package main

import (
	"log"
	"net/http"

	"github.com/algoristas/api/problems"
	"github.com/algoristas/api/results"
	"github.com/algoristas/api/router"
	"github.com/algoristas/api/standings"
	"github.com/algoristas/api/users"
)

func main() {
	log.Println("Listening at :8080...")
	http.ListenAndServe(":8080", router.Wire(router.Dependencies{
		StandingsDataProvider: standings.NewDataProvider(),
		ResultsDataProvider:   results.NewDataProvider(),
		ProblemsDataProvider:  problems.NewDataProvider(),
		UsersDataProvider:     users.NewDataProvider(),
	}))
}

package main

import (
	"log"
	"net/http"

	"github.com/algoristas/api/problems"
	"github.com/algoristas/api/results"
	"github.com/algoristas/api/router"
	"github.com/algoristas/api/standings"
)

func main() {
	log.Println("Listening at :8080...")
	http.ListenAndServe(":8080", router.Wire(router.Dependencies{
		StandingsDAO: standings.NewStandingsDAO(),
		ResultsDAO:   results.NewResultsDAO(),
		ProblemsDAO:  problems.NewProblemsDAO(),
	}))
}

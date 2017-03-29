package main

import (
	"net/http"

	"github.com/algoristas/api/controllers"
)

func main() {
	standingsController := controllers.NewStandingsController()
	http.HandleFunc("/v1/standings", standingsController.Index)

	resultsController := controllers.NewResultsController()
	http.HandleFunc("/v1/results", resultsController.Index)

	http.ListenAndServe(":8080", nil)
}

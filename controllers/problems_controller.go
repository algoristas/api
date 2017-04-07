package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/algoristas/api/problems"
)

type ProblemsController struct {
}

func (t *ProblemsController) SetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := problems.GetSets()
	if err != nil {
		log.Printf("Failed to retrieve results: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func NewProblemsController() *ProblemsController {
	return &ProblemsController{}
}

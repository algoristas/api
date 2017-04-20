package problems

import (
	"fmt"
	"log"
	"net/http"
)

// ProblemsController describes controller for requests at /problems/.
type ProblemsController struct {
	dao DAO
}

// SetIndex handles /problems/sets endpoint, returns all problems.
func (t *ProblemsController) SetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := t.dao.GetSets()
	if err != nil {
		log.Printf("Failed to retrieve results: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// NewController returns a new initialized ProblemsController.
func NewController(problemsDAO DAO) *ProblemsController {
	return &ProblemsController{
		dao: problemsDAO,
	}
}

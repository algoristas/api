package problems

import (
	"fmt"
	"log"
	"net/http"

	"github.com/algoristas/api/users"
	"github.com/julienschmidt/httprouter"
	"github.com/rendon/httpresp"
	"strconv"
)

// Controller describes controller for requests at /problems/.
type Controller struct {
	problemsDataProvider DataProvider
	usersDataProvider    users.DataProvider
}

// SetIndex handles /problems/sets endpoint, returns all problems.
func (t *Controller) SetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data, err := t.problemsDataProvider.GetSets()
	if err != nil {
		log.Printf("Failed to retrieve results: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetProblem handles the /users/:userID/problems/:problemID endpoint, returns problem details
// for a given user.
func (t *Controller) GetProblem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	userID := params.ByName("userID")
	user, err := t.usersDataProvider.FindUser(userID)
	if err != nil {
		httpresp.Error(w, "User not found", http.StatusNotFound)
		return
	}

	problemID, err := strconv.Atoi(params.ByName("problemID"))
	if err != nil {
		httpresp.BadRequest(w, "Invalid problem ID")
	}

	problem, err := t.problemsDataProvider.FindProblem(user.UserName, uint(problemID))
	if err != nil {
		log.Printf("Failed to retrieve problem for %s/%d: %s", userID, problemID, err)
		httpresp.ServerError(w, "Failed to retrieve problem")
		return
	}

	httpresp.Data(w, problem, http.StatusOK)
}

// NewController returns a new initialized Controller.
func NewController(problemsDataProvider DataProvider, usersDataProvider users.DataProvider) *Controller {
	return &Controller{
		problemsDataProvider: problemsDataProvider,
		usersDataProvider:    usersDataProvider,
	}
}

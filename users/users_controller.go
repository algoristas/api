package users

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

// Controller describes controller for requests at /users/.
type Controller struct {
	usersDataProvider DataProvider
}

// GetUser handles the /users/:userId endpoint.
func (t *Controller) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	userID := params.ByName("userId")

	user, userErr := t.usersDataProvider.FindUser(userID)
	if userErr != nil {
		if userErr == ErrNotFound {
			log.Printf("User not found: %s", userID)
			userNotFound(w, "User not found")
			return
		}
		log.Printf("Failed to retrieve user: %s", userErr)
		internalServerError(w, "Internal Server Error")
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Printf("Failed to serialize response: %s", err)
		internalServerError(w, "Failed to generate response")
		return
	}
	fmt.Fprintf(w, "%s", data)
}

// NewController returns an initialized Controller.
func NewController(userDataProvider DataProvider) *Controller {
	return &Controller{
		usersDataProvider: userDataProvider,
	}
}

func internalServerError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, `{"status": 500, "message": "%s"}`, message)
}

func userNotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"status": 404, "message": "%s"}`, message)
}

func badRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"status": 400, "message": "%s"}`, message)
}

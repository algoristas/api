package users

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"encoding/json"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

var numberRegexp = regexp.MustCompile(`^\d+$`)

// Controller describes controller for requests at /users/.
type Controller struct {
	usersDataProvider DataProvider
}

// GetUser handles the /users/:userId endpoint.
func (t *Controller) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	userID := params.ByName("userId")

	var user *User
	var userErr error

	if numberRegexp.MatchString(userID) {
		id, err := strconv.ParseInt(userID, 10, 32)
		if err != nil {
			log.Printf("Invalid ID (%s): %s", userID, err)
			badRequest(w, "Invalid ID")
			return
		}
		user, userErr = t.usersDataProvider.FindUserByID(uint(id))
	} else {
		user, userErr = t.usersDataProvider.FindUserByUserName(userID)
	}

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

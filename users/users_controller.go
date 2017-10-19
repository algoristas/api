package users

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"github.com/rendon/httpresp"
)

// Controller describes controller for requests at /users/.
type Controller struct {
	usersDataProvider DataProvider
}

// GetUser handles the /users/:userID endpoint.
func (t *Controller) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	userID := params.ByName("userID")

	user, userErr := t.usersDataProvider.FindUser(userID)
	if userErr != nil {
		if userErr == ErrNotFound {
			log.Printf("User not found: %s", userID)
			httpresp.NotFound(w, "User not found")
			return
		}
		log.Printf("Failed to retrieve user: %s", userErr)
		httpresp.ServerError(w, "Internal Server Error")
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Printf("Failed to serialize response: %s", err)
		httpresp.ServerError(w, "Failed to generate response")
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

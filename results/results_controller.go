package results

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Controller describes controller for requests at /results/.
type Controller struct {
	dataProvider DataProvider
}

// Index handles /results/ endpoint, returns all results.
func (t *Controller) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data, err := t.dataProvider.GetResults()
	if err != nil {
		log.Printf("Failed to retrieve results: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": 500, "message": "Failed to retrieve data"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// NewController returns an initialized Controller.
func NewController(dataProvider DataProvider) *Controller {
	return &Controller{
		dataProvider: dataProvider,
	}
}

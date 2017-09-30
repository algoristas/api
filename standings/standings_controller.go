package standings

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Controller describes controller for requests at /standings/.
type Controller struct {
	dataProvider DataProvider
}

// Index handles /standings endpoint, returns the list of standings.
func (t *Controller) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data, err := t.dataProvider.GetStandings()
	if err != nil {
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

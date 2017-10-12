package test

import (
	"encoding/json"
	"net/http"

	. "github.com/onsi/gomega"

	"github.com/algoristas/api/model"
)

// ReadError reads data from response body and maps that content to the provided interface.
func ReadError(r *http.Response) model.ErrorResponse {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var errorResponse model.ErrorResponse
	Expect(decoder.Decode(&errorResponse)).To(BeNil())
	return errorResponse
}

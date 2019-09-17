package stats

import (
	"net/http"

	"github.com/amnay-mo/fizzbuzz-api/utils"
)

// HTTPHandler is the http handler for the stats endpoint
type HTTPHandler struct {
	Store interface {
		GetMax() (*Stats, error)
	}
}

// ServeHTTP returns the set of parameters with most hits and its respective hit count
func (a HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s, err := a.Store.GetMax()
	if err != nil {
		utils.Jsonify(w, &errorBody{Error: err.Error()}, http.StatusInternalServerError)
		return
	}
	utils.Jsonify(w, s, http.StatusOK)
}

type errorBody struct {
	Error string `json:"error"`
}

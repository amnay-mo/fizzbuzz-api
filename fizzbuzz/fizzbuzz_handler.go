package fizzbuzz

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amnay-mo/fizzbuzz-api/utils"
)

// Body is the body returned by the API
type Body struct {
	Sequence []string `json:"sequence"`
}

// ErrorBody is the body returned by the API
type ErrorBody struct {
	Error string `json:"error"`
}

// nolint:gocyclo
func parseFizzBuzzParameters(r *http.Request) (*Parameters, error) {
	fizzNumberStr := r.URL.Query().Get("fizzNumber")
	if fizzNumberStr == "" {
		return nil, fmt.Errorf("fizzNumber missing in query parameters")
	}
	fizzNumber, err := strconv.Atoi(fizzNumberStr)
	if err != nil || fizzNumber < 0 {
		return nil, fmt.Errorf("invalid fizzNumber: %v", fizzNumberStr)
	}
	buzzNumberStr := r.URL.Query().Get("buzzNumber")
	if buzzNumberStr == "" {
		return nil, fmt.Errorf("buzzNumber missing in query parameters")
	}
	buzzNumber, err := strconv.Atoi(buzzNumberStr)
	if err != nil || buzzNumber < 0 {
		return nil, fmt.Errorf("invalid buzzNumber: %v", buzzNumberStr)
	}
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		return nil, fmt.Errorf("limit missing in query parameters")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		return nil, fmt.Errorf("invalid limit: %v", limitStr)
	}
	fizzWord := r.URL.Query().Get("fizzWord")
	if fizzWord == "" {
		return nil, fmt.Errorf("fizzWord missing in query parameters")
	}
	buzzWord := r.URL.Query().Get("buzzWord")
	if buzzWord == "" {
		return nil, fmt.Errorf("buzzWord missing in query parameters")
	}
	return &Parameters{
			FizzNumber: fizzNumber,
			BuzzNumber: buzzNumber,
			Limit:      limit,
			FizzWord:   fizzWord,
			BuzzWord:   buzzWord},
		nil
}

// HandleFizzBuzz returns a FizzBuzz sequence according to parameters
func HandleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	fbb, err := parseFizzBuzzParameters(r)
	if err != nil {
		utils.Jsonify(w, &ErrorBody{err.Error()}, http.StatusBadRequest)
		return
	}
	sequence := Sequence(fbb)
	body := Body{Sequence: sequence}
	utils.Jsonify(w, &body, http.StatusOK)
}

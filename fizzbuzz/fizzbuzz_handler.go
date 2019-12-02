package fizzbuzz

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amnay-mo/fizzbuzz-api/utils"
	"github.com/pkg/errors"
)

// Body is the body returned by the API
type Body struct {
	Sequence []string `json:"sequence"`
}

// ErrorBody is the body returned by the API
type ErrorBody struct {
	Error string `json:"error"`
}

var ErrInvalidParameters = errors.New("invalid fizzbuzz parameters")

// nolint:gocyclo
func parseFizzBuzzParameters(r *http.Request) (*Parameters, error) {
	fizzNumberStr := r.URL.Query().Get("fizzNumber")
	if fizzNumberStr == "" {
		return nil, errors.Wrap(ErrInvalidParameters, "fizzNumber missing in query parameters")
	}
	fizzNumber, err := strconv.Atoi(fizzNumberStr)
	if err != nil || fizzNumber < 0 {
		return nil, errors.Wrap(ErrInvalidParameters, fmt.Sprintf("invalid fizzNumber: %v", fizzNumberStr))
	}
	buzzNumberStr := r.URL.Query().Get("buzzNumber")
	if buzzNumberStr == "" {
		return nil, errors.Wrap(ErrInvalidParameters, "buzzNumber missing in query parameters")
	}
	buzzNumber, err := strconv.Atoi(buzzNumberStr)
	if err != nil || buzzNumber < 0 {
		return nil, errors.Wrap(ErrInvalidParameters, fmt.Sprintf("invalid buzzNumber: %v", buzzNumberStr))
	}
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		return nil, errors.Wrap(ErrInvalidParameters, "limit missing in query parameters")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		return nil, errors.Wrap(ErrInvalidParameters, fmt.Sprintf("invalid limit: %v", limitStr))
	}
	fizzWord := r.URL.Query().Get("fizzWord")
	if fizzWord == "" {
		return nil, errors.Wrap(ErrInvalidParameters, fmt.Sprintf("fizzWord missing in query parameters"))
	}
	buzzWord := r.URL.Query().Get("buzzWord")
	if buzzWord == "" {
		return nil, errors.Wrap(ErrInvalidParameters, "buzzWord missing in query parameters")
	}
	return &Parameters{
			FizzNumber: fizzNumber,
			BuzzNumber: buzzNumber,
			Limit:      limit,
			FizzWord:   fizzWord,
			BuzzWord:   buzzWord},
		nil
}

// GetHandlerFunc returns a FizzBuzz handler
func GetHandlerFunc(f SequenceFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fbb, err := parseFizzBuzzParameters(r)
		if err != nil {
			handleErr(w, err)
			return
		}
		sequence, err := f(fbb)
		if err != nil {
			handleErr(w, err)
			return
		}
		body := Body{Sequence: sequence}
		utils.Jsonify(w, &body, http.StatusOK)
	}
}

func handleErr(w http.ResponseWriter, err error) {
	switch errors.Cause(err) {
	case ErrInvalidParameters:
		utils.Jsonify(w, &ErrorBody{err.Error()}, http.StatusBadRequest)
	case ErrMaxLimitExceeded:
		utils.Jsonify(w, &ErrorBody{err.Error()}, http.StatusPreconditionFailed)
	default:
		utils.Jsonify(w, &ErrorBody{err.Error()}, http.StatusInternalServerError)
	}
}

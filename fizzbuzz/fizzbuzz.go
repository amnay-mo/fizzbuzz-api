package fizzbuzz

import (
	"errors"
	"strconv"
)

func isMultiple(lhn, rhn int) bool {
	switch {
	case rhn == 0:
		return lhn == 0
	case lhn == 0:
		return true
	default:
		return lhn%rhn == 0
	}
}

// Parameters are the parameters used to produce the FizzBuzz Sequence
type Parameters struct {
	FizzNumber, BuzzNumber, Limit int
	FizzWord, BuzzWord            string
}

// SequenceFunc is a fizzbuzz sequence function
type SequenceFunc func(params *Parameters) ([]string, error)

// WithMaxLimit adds max limit check on a fizzbuzz sequence function
func WithMaxLimit(f SequenceFunc, maxLimit int) SequenceFunc {
	if maxLimit < 1 {
		return f
	}
	return func(params *Parameters) ([]string, error) {
		if params.Limit > maxLimit {
			return nil, ErrMaxLimitExceeded
		}
		return f(params)
	}
}

// Sequence returns a fizzbuzz sequence based on the input
func Sequence(params *Parameters) ([]string, error) {
	fizzBuzzWord := params.FizzWord + params.BuzzWord
	sequence := make([]string, params.Limit)
	var number int
	for i := 0; i < params.Limit; i++ {
		number = i + 1
		switch {
		case isMultiple(number, params.FizzNumber) && isMultiple(number, params.BuzzNumber):
			sequence[i] = fizzBuzzWord
		case isMultiple(number, params.FizzNumber):
			sequence[i] = params.FizzWord
		case isMultiple(number, params.BuzzNumber):
			sequence[i] = params.BuzzWord
		default:
			sequence[i] = strconv.Itoa(number)
		}
	}
	return sequence, nil
}

// ErrMaxLimitExceeded is an error thrown when the max limit is exceeded
var ErrMaxLimitExceeded = errors.New("max limit exceeded")

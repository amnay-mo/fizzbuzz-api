package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amnay-mo/fizzbuzz-api/utils"
)

func TestHandleFizzBuzz(t *testing.T) {
	tt := []struct {
		url      string
		sequence []string
		status   int
	}{
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=3&limit=10&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{"1", "Bonnie", "Clyde", "Bonnie", "5", "BonnieClyde", "7", "Bonnie", "Clyde", "Bonnie"},
			http.StatusOK,
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=-1&buzzNumber=3&limit=10&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=notanumber&limit=10&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=3&limit=infinite&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=3&limit=10",
			[]string{},
			http.StatusBadRequest,
		},
		{
			"localhost:8080/fizzbuzz?fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
		},
	}
	for _, tc := range tt {
		req, err := http.NewRequest(
			"GET",
			tc.url,
			nil,
		)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		HandleFizzBuzz(rec, req)
		res := rec.Result()
		if res.StatusCode != tc.status {
			t.Errorf("Expected Status %v, got: %v", tc.status, res.StatusCode)
		}
		if tc.status == http.StatusOK {
			fbb := new(FizzBuzzBody)
			dec := json.NewDecoder(res.Body)
			err = dec.Decode(fbb)
			if err != nil {
				t.Fatalf("Could not decode body: %v", err)
			}
			expectedSequence := []string{"1", "Bonnie", "Clyde", "Bonnie", "5", "BonnieClyde", "7", "Bonnie", "Clyde", "Bonnie"}
			if !utils.AreEqualStringSlices(fbb.Sequence, expectedSequence) {
				t.Errorf("Expected: %v, got %v", expectedSequence, fbb.Sequence)
			}
		}
	}
}

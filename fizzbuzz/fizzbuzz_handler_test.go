package fizzbuzz

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestHandleFizzBuzz(t *testing.T) {
	tt := []struct {
		url         string
		sequence    []string
		status      int
		errorPrefix string
	}{
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=3&limit=10&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{"1", "Bonnie", "Clyde", "Bonnie", "5", "BonnieClyde", "7", "Bonnie", "Clyde", "Bonnie"},
			http.StatusOK,
			"",
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=-1&buzzNumber=3&limit=10&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
			"invalid fizzNumber",
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=notanumber&limit=10&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
			"invalid buzzNumber",
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=3&limit=infinite&fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
			"invalid limit",
		},
		{
			"localhost:8080/fizzbuzz?fizzNumber=2&buzzNumber=3&limit=10",
			[]string{},
			http.StatusBadRequest,
			"fizzWord missing",
		},
		{
			"localhost:8080/fizzbuzz?fizzWord=Bonnie&buzzWord=Clyde",
			[]string{},
			http.StatusBadRequest,
			"fizzNumber missing",
		},
	}
	h := GetHandlerFunc(Sequence)
	for _, tc := range tt {
		req, err := http.NewRequest(
			"GET",
			tc.url,
			nil,
		)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		h(rec, req)
		res := rec.Result()
		if res.StatusCode != tc.status {
			t.Errorf("expected status %v, got: %v", tc.status, res.StatusCode)
		}
		if tc.status == http.StatusOK {
			fbb := new(Body)
			dec := json.NewDecoder(res.Body)
			err = dec.Decode(fbb)
			if err != nil {
				t.Fatalf("could not decode body: %v", err)
			}
			if !reflect.DeepEqual(tc.sequence, fbb.Sequence) {
				t.Errorf("expected sequence: %v, got: %v", tc.sequence, fbb.Sequence)
			}
		} else {
			fbeb := new(ErrorBody)
			dec := json.NewDecoder(res.Body)
			err = dec.Decode(fbeb)
			if err != nil {
				t.Fatalf("could not decode body: %v", err)
			}
			if !strings.HasPrefix(fbeb.Error, tc.errorPrefix) {
				t.Errorf("expected error message starting with: %s, got: %s", tc.errorPrefix, fbeb.Error)
			}
		}
	}
}

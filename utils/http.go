package utils

import (
	"encoding/json"
	"net/http"
)

// Jsonify does the trick!
func Jsonify(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	enc.Encode(data) // nolint:errcheck
}

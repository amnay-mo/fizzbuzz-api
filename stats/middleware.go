package stats

import (
	"net/http"
)

// Middleware keeps track of the most hit HTTP query parameters
type Middleware struct {
	Next  http.Handler
	Store interface {
		Increment(string) error
	}
}

// ServeHTTP is the stats middleware's handler function
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Next.ServeHTTP(w, r)
	// Encode() will always sort the query parameters in ascending order
	// This is useful as we track hit count regardless of parameter order
	m.Store.Increment(r.URL.Query().Encode())
}

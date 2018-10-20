package utils

import (
	"log"
	"net/http"
)

// LoggerMiddleware is a logging middleware
type LoggerMiddleware struct {
	Next http.Handler
}

// Handle is LoggerMiddleware's implementation
func (l LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.Next.ServeHTTP(w, r)
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

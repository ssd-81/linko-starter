package main

import (
	"net/http"
)

func (s *server) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		logger.Printf("Served request: %v %v", r.Method, r.URL.Path)
	})
}

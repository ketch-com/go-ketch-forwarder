package server

import (
	"net/http"
)

// NewServer returns a new http.Server
func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Handler: handler,
	}
}

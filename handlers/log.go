package handlers

import (
	"net/http"
	"log"
)

// AccessLogger logs access
func AccessLogger(h http.Handler) http.Handler {
	return &accessLogger{h}
}

type accessLogger struct {
	Handler http.Handler
}

func (a *accessLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Access:", r.Method, ":", r.URL)
	a.Handler.ServeHTTP(w, r)
}

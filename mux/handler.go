package mux

import (
	"errors"
	"html"
	"log"
	"net/http"

	"github.com/mock-http-server/config"
)

func Must(vars config.Handler) http.Handler {

	mockEndpoints, err := loadFromDirectory(vars.Directory)
	if err != nil {
		log.Panic(err)
	}
	m := http.NewServeMux()
	m.HandleFunc("/", newHandler(mockEndpoints))
	return m
}

func newHandler(mockEndpoints []*MockEndpoint) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Request, %q\r\n", html.EscapeString(r.URL.Path))

		endpoint, err := getMockEndpoint(r, mockEndpoints)
		var mockErr *Error
		switch {
		case errors.As(err, &mockErr):
			mockErr.Send(w)
			return
		case err != nil:
			log.Panic(err)
		default:
		}

		endpoint.Response.Send(w)
	}
}

package mux

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/mock-http-server/config"
)

func Must(vars config.Handler) http.Handler {

	mockHandlers, err := loadFromDirectory(vars.Directory)
	if err != nil {
		log.Panic(err)
	}
	m := http.NewServeMux()
	m.HandleFunc("/", newHandler(mockHandlers))
	return m
}

func newHandler(mockHandlers []*MockHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := getMockHandler(r, mockHandlers)
		var mockErr *Error
		switch {
		case errors.As(err, &mockErr):
			mockErr.Send(w)
			return
		case err != nil:
			log.Panic(err)
		default:
		}
		fmt.Fprintf(w, "Request, %q\r\n", html.EscapeString(r.URL.Path))
	}
}

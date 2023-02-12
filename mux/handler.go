package mux

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/mock-http-server/config"
)

func Must(vars config.Handler) http.Handler {

	_, err := loadFromDirectory(vars.Directory)
	if err != nil {
		log.Panic(err)
	}
	m := http.NewServeMux()
	m.HandleFunc("/", newHandler())
	return m
}

func newHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Request, %q", html.EscapeString(r.URL.Path))
	}
}

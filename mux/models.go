package mux

import (
	"encoding/json"
	"log"
	"net/http"
)

type MockHandler struct {
	Path   Path   `json:"path"`
	Method string `json:"method"`
}

type Path struct {
	Pattern   string     `json:"pattern"`
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Label string `json:"name"`
	Value string `json:"value"`
}

type Error struct {
	statusCode int
	Message    string `json:"message"`
	Method     string `json:"method"`
	Path       string `json:"path"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Send(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		log.Printf("unable to write http error %v\n", err)
		return
	}
}

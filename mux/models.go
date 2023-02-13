package mux

import (
	"encoding/json"
	"log"
	"net/http"
)

type MockEndpoint struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
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

type Response struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
}

func (r *Response) Send(w http.ResponseWriter) {

	if r.Body != nil {
		w.Header().Add("Content-Type", "application/json")
	}
	w.WriteHeader(r.StatusCode)

	if r.Body != nil {
		if err := json.NewEncoder(w).Encode(r.Body); err != nil {
			log.Printf("unable to write http error %v\n", err)
			return
		}
	}
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

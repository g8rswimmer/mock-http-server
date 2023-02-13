package mux

import (
	"net/http"
	"strings"
)

func getMockHandler(req *http.Request, mockHandlers []*MockHandler) (*MockHandler, error) {
	p := req.URL.Path
	handlers := map[string]*MockHandler{}

	for _, mh := range mockHandlers {
		if mh.Path.Pattern == p {
			method := strings.ToUpper(mh.Method)
			if method == "" {
				method = http.MethodGet
			}
			handlers[method] = mh
		}
	}
	h, ok := handlers[req.Method]
	if !ok {
		return nil, &Error{
			statusCode: http.StatusNotFound,
			Message:    "unable to retrieve mock handler",
			Path:       req.URL.Path,
			Method:     req.Method,
		}
	}
	return h, nil
}

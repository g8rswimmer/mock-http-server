package mux

import (
	"net/http"
	"strings"
)

func getMockEndpoint(req *http.Request, mockEndpoints []*MockEndpoint) (*MockEndpoint, error) {
	p := req.URL.Path
	endpoints := map[string]*MockEndpoint{}

	for _, mh := range mockEndpoints {
		if mh.Request.Path.Pattern == p {
			method := strings.ToUpper(mh.Request.Method)
			if method == "" {
				method = http.MethodGet
			}
			endpoints[method] = mh
		}
	}
	h, ok := endpoints[req.Method]
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

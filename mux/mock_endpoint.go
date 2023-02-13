package mux

import (
	"net/http"
	"strings"
)

func getMockEndpoint(req *http.Request, mockEndpoints []*MockEndpoint) (*MockEndpoint, error) {
	p := req.URL.Path
	endpoints := map[string]*MockEndpoint{}

	for _, mp := range mockEndpoints {
		if err := comparePath(p, mp.Request.Path); err == nil {
			method := strings.ToUpper(mp.Request.Method)
			if method == "" {
				method = http.MethodGet
			}
			endpoints[method] = mp
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

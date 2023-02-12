package mux

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

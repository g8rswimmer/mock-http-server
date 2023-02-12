package mux

type Handler struct {
	Path     Path        `json:"path"`
	Method   string      `json:"method"`
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
}

type Path struct {
	Pattern   string     `json:"pattern"`
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Label string `json:"name"`
	Value string `json:"value"`
}

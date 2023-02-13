package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

const (
	port        = "MOCK_HTTP_SRVR_PORT"
	defaultPort = "8080"
	directory   = "MOCK_HTTP_HANDLER_DIR"
)

type Server struct {
	Port string
}

type Handler struct {
	Directory string
}

type Vars struct {
	Server  Server
	Handler Handler
}

func MustLoad() *Vars {
	v := &Vars{
		Server:  mustServer(),
		Handler: mustHandler(),
	}

	enc, err := json.MarshalIndent(v, "", "    ")
	if err == nil {
		log.Println("Config Vars...")
		log.Println(string(enc))
	}
	return v
}

func mustServer() Server {
	return Server{
		Port: mustPort(),
	}
}

func mustPort() string {
	p := os.Getenv(port)
	if p == "" {
		return defaultPort
	}
	_, err := strconv.Atoi(p)
	if err != nil {
		log.Panicf("%s not defined as an init %v", port, err)
	}
	return p
}

func mustHandler() Handler {
	return Handler{
		Directory: os.Getenv(directory),
	}
}

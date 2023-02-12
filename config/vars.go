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
)

type Server struct {
	Port string
}
type Vars struct {
	Server Server
}

func MustLoad() *Vars {
	v := &Vars{
		Server: Server{
			Port: mustPort(),
		},
	}

	enc, err := json.MarshalIndent(v, "", "    ")
	if err == nil {
		log.Println("Config Vars...")
		log.Println(string(enc))
	}
	return v
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

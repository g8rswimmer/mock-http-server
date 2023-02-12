package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

const (
	port        = "MOCK_HTTP_PORT"
	defaultPort = "8080"
)

type Vars struct {
	Port string
}

func MustLoad() *Vars {
	v := &Vars{
		Port: mustPort(),
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
		log.Panicf("MOCK_HTTP_PORT not defined as an init %v", err)
	}
	return p
}

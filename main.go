package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mock-http-server/config"
	"github.com/mock-http-server/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Starting Mock HTTP Server...")
	vars := config.MustLoad()

	log.Println("Happy Mocking...")
	srvrShutdown := server.Start(vars)
	defer srvrShutdown(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Mock HTTP Server shutting down...")
}

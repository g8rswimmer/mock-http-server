package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mock-http-server/config"
	"github.com/mock-http-server/mux"
)

type Shutdown func(context.Context)

func Start(vars config.Server) Shutdown {

	h := mux.New()

	srvr := &http.Server{
		Addr:    ":" + vars.Port,
		Handler: h,
	}

	go func() {
		log.Println("Starting Server...")
		err := srvr.ListenAndServe()
		switch {
		case errors.Is(err, http.ErrServerClosed):
		case err != nil:
			log.Panicf("server error %v", err)
		default:
		}
	}()

	return func(ctx context.Context) {
		sCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		if err := srvr.Shutdown(sCtx); err != nil {
			log.Panicf("server shutdown error %v", err)
		}
	}
}

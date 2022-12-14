package main

import (
	"context"
	"cooks-assistant/internal/recipes/ports"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	_ = context.Background()
	apiRouter := chi.NewRouter()
	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", setup(apiRouter))

	err := http.ListenAndServe(":6000", rootRouter)
	if err != nil {
		log.Fatalf("unable to start http server: %s", err)
	}
}

func setup(router chi.Router) http.Handler {
	return ports.HandlerFromMux(ports.NewHttpServer(), router)
}

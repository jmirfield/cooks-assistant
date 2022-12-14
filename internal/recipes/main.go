package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jmirfield/cooks-assistant/internal/common/auth"
	"github.com/jmirfield/cooks-assistant/internal/recipes/ports"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
)

func main() {
	_ = context.Background()
	apiRouter := chi.NewRouter()
	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", setup(apiRouter))

	config := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT")}
	firebaseApp, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatal("Unable to create firebase Auth client")
	}

	apiRouter.Use(auth.FirebaseHttpMiddleware{authClient}.Middleware)

	err = http.ListenAndServe(":6000", rootRouter)
	if err != nil {
		log.Fatalf("unable to start http server: %s", err)
	}
}

func setup(router chi.Router) http.Handler {
	return ports.HandlerFromMux(ports.NewHttpServer(), router)
}

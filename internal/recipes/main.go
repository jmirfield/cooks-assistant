package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jmirfield/cooks-assistant/internal/common/auth"
	"github.com/jmirfield/cooks-assistant/internal/recipes/ports"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
)

func main() {
	_ = context.Background()
	apiRouter := chi.NewRouter()

	var opts []option.ClientOption
	opts = append(opts, option.WithCredentialsFile("./cooks-assistant-dev-firebase-adminsdk-ough3-abf67b2303.json"))

	config := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT")}
	firebaseApp, err := firebase.NewApp(context.Background(), config, opts...)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatal("Unable to create firebase Auth client: %s", err)
	}

	apiRouter.Use(auth.FirebaseHttpMiddleware{authClient}.Middleware)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", setup(apiRouter))
	err = http.ListenAndServe(":6000", rootRouter)
	if err != nil {
		log.Fatalf("unable to start http server: %s", err)
	}
}

func setup(router chi.Router) http.Handler {
	return ports.HandlerFromMux(ports.NewHttpServer(), router)
}

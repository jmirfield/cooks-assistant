package auth

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/go-chi/render"
)

type FirebaseHttpMiddleware struct {
	AuthClient *auth.Client
}

func (a FirebaseHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			if err := render.Render(w, r, nil); err != nil {
				panic(err)
			}
		}

		token, err := a.AuthClient.VerifyIDToken(ctx, bearerToken)
		if err != nil {
			if err := render.Render(w, r, nil); err != nil {
				panic(err)
			}
		}

		ctx = context.WithValue(ctx, "id", token.UID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (a FirebaseHttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}

package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/code-wave/go-wave/infrastructure/auth"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type contextKey string

var ContextKeyTokenUserID = contextKey("user_id")

//case 1: access token by payload
func AuthVerifyPayloadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.SetJsonHeader(w)
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			err := errors.NewUnauthorizedError("no Authorization header provided")
			w.WriteHeader(err.Status)
			w.Write(err.ResponseJSON().([]byte))
			return
		}

		clientToken := auth.ExtractToken(bearerToken)
		if clientToken == "" {
			err := errors.NewUnauthorizedError("invalid format of authorization header")
			w.WriteHeader(err.Status)
			w.Write(err.ResponseJSON().([]byte))
			return
		}

		claims, err := auth.JwtWrapper.ValidateToken(clientToken)
		if err != nil {
			authErr := errors.NewUnauthorizedError(err.Error())
			log.Println(authErr)
			w.WriteHeader(authErr.Status)
			w.Write(authErr.ResponseJSON().([]byte))
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyTokenUserID, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//case 2: access token by cookie
func AuthVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.SetJsonHeader(w)

		atCookie, err := r.Cookie("access_token")
		if err != nil {
			log.Println("access token from client's cookie doesn't exist " + err.Error())
			authErr := errors.NewUnauthorizedError(err.Error())
			w.WriteHeader(authErr.Status)
			w.Write(authErr.ResponseJSON().([]byte))
			return
		}

		claims, err := auth.JwtWrapper.ValidateToken(atCookie.Value)
		if err != nil {
			authErr := errors.NewUnauthorizedError(err.Error())
			log.Println(authErr)
			w.WriteHeader(authErr.Status)
			w.Write(authErr.ResponseJSON().([]byte))
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyTokenUserID, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

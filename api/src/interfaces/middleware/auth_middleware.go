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

func AuthVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.SetJsonHeader(w)
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			err := errors.NewForbiddenError("no Authorization header provided")
			w.WriteHeader(err.Status)
			w.Write(err.ResponseJSON().([]byte))
			return
		}

		clientToken := auth.ExtractToken(bearerToken)
		if clientToken == "" {
			err := errors.NewBadRequestError("invalid format of authorization header")
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

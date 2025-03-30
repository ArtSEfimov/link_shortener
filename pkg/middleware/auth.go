package middleware

import (
	"context"
	"http_server/configs"
	"http_server/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	auth := func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer") {
			writeUnauthorized(w)
			return
		}
		token := strings.Fields(header)[1]

		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthorized(w)
			return
		}

		newContext := context.WithValue(r.Context(), ContextEmailKey, data.Email)

		newRequest := r.WithContext(newContext)

		next.ServeHTTP(w, newRequest)
	}
	return http.HandlerFunc(auth)
}

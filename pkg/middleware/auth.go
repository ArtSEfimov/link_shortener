package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	auth := func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		token := strings.Fields(header)[1]
		fmt.Println(token)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(auth)
}

package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	logger := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapper := &WriterWrapper{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		log.Println(wrapper.StatusCode, r.URL.Path, r.Method, time.Since(start))

	}
	return http.HandlerFunc(logger)
}

package middleware

import "net/http"

type WriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WriterWrapper) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

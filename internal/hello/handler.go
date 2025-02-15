package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func NewHalloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (h *HelloHandler) Hello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello World")
	}
}

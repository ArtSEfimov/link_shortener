package main

import (
	"fmt"
	"http_server/configs"
	"http_server/internal/hello"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)

	router := http.NewServeMux()

	hello.NewHalloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	listenErr := server.ListenAndServe()
	if listenErr != nil {
		panic(listenErr)
	}
}

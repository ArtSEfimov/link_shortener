package main

import (
	"fmt"
	"http_server/configs"
	"http_server/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	listenErr := server.ListenAndServe()
	if listenErr != nil {
		panic(listenErr)
	}
}

package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	db.NewDB(conf) // TODO: there returns *db
	fmt.Println("Listening...")
	//fmt.Println(conf)

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

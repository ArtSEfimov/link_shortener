package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/internal/link"
	"http_server/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	dataBase := db.NewDB(conf)
	fmt.Println("Listening...")
	//fmt.Println(conf)

	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewRepository(dataBase)

	// Handler
	auth.NewHandler(router, auth.HandlerDeps{
		Config: conf,
	})

	link.NewHandler(router, link.HandlerDeps{LinkRepository: linkRepository})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	listenErr := server.ListenAndServe()
	if listenErr != nil {
		panic(listenErr)
	}
}

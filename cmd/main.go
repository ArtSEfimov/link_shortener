package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/internal/link"
	"http_server/internal/statistic"
	"http_server/internal/user"
	"http_server/pkg/db"
	"http_server/pkg/event"
	"http_server/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	dataBase := db.NewDB(conf)
	fmt.Println("Listening...")

	router := http.NewServeMux()

	eventBus := event.NewBus()

	// Repositories
	linkRepository := link.NewRepository(dataBase)
	userRepository := user.NewRepository(dataBase)
	statisticRepository := statistic.NewRepository(dataBase)

	// Service
	authService := auth.NewService(userRepository)
	statisticService := statistic.NewService(statistic.ServiceDeps{
		EventBus:            eventBus,
		StatisticRepository: statisticRepository,
	})

	// Handler
	auth.NewHandler(router, auth.HandlerDeps{
		Config:  conf,
		Service: authService,
	})

	link.NewHandler(router, link.HandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})

	statistic.NewHandler(router, statistic.HandlerDeps{
		Config:              conf,
		StatisticRepository: statisticRepository,
	})

	// Middlewares
	middlewares := middleware.Chain(
		middleware.Logging,
		middleware.CORS,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: middlewares(router),
	}

	go statisticService.AddClick()

	listenErr := server.ListenAndServe()
	if listenErr != nil {
		panic(listenErr)
	}
}

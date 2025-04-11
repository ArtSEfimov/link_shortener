package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"http_server/internal/link"
	"http_server/internal/statistic"
	"http_server/internal/user"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, openErr := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if openErr != nil {
		panic(openErr)
	}

	autoMigrateErr := db.AutoMigrate(&link.Link{}, &user.User{}, &statistic.Statistic{})
	if autoMigrateErr != nil {
		return
	}
}

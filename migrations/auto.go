package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"http_server/internal/link"
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

	autoMigrateErr := db.AutoMigrate(&link.Link{})
	if autoMigrateErr != nil {
		return
	}
}

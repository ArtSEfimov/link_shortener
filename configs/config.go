package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Panicln("Error loading .env file")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}

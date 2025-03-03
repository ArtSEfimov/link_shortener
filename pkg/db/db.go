package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"http_server/configs"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *DB {
	db, openErr := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if openErr != nil {
		panic(openErr)
	}
	return &DB{db}
}

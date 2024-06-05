package app

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DBConnection() *gorm.DB {
	DB_POSTGRES_URL := os.Getenv("DB_POSTGRES_URL")
	dsn := DB_POSTGRES_URL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	Db = db
	return db
}

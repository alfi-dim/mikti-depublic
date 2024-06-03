package app

import (
	"log"
	"mikti-depublic/model/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	dsn := os.Getenv("DB_POSTGRES_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = DB.AutoMigrate(&domain.User{}, &domain.Admin{}, &domain.Event{}, &domain.Transaction{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}

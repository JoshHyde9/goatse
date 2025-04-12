package db

import (
	"fmt"
	"log"

	"goatse/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialise() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=goatse port=5432 sslmode=disable TimeZone=Australia/Sydney"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
		return nil, err
	}

	fmt.Println("Successfully connected to db!")

	db.AutoMigrate(&models.Post{})

	return db, nil
}

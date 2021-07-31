package database

import (
	// gorm dialect
	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/dy-fi/war-room/models"
)

// Connect to database and return client object
func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=dev dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// db is needed so panic out
		panic(err)
	}

	// All models here 
	db.AutoMigrate(&models.City{}, &models.Place{})
	return db
}
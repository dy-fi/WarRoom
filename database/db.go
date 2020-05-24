package database

import (
	// gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"

	"github.com/dy-fi/war-room/models"
)

// Connect to database and return client object
func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=postgres user=dev password=dev dbname=wrdb")
	defer db.Close()

	if err != nil {
		// db is needed so panic out
		panic(err)
	}

	// All models here 
	db.AutoMigrate(&models.City{}, &models.Place{})
	return db
}
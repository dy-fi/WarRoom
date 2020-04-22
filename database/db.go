package database

import (
	// gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"

)

// Connect to database and return client object
func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=dev password=dev dbname=wrdb")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	
	return db
}
package db

import (
	"fmt"

	"github.com/dy-fi/war-room/models"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Connect to MySql db
func Connect() *gorm.DB {
	// parseTime to convert gorm createdAt ISO to sql created_at uint8[]
	db, err := gorm.Open("mysql", "WarRoom:123456@tcp(:3333)/wrdb?parseTime=true")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connected succesfully to db")
	}
	db.LogMode(true)
	// defer db.Close() // turn on for testing

	// Turn off in production, turn on in dev
	db.AutoMigrate(&models.User{}, &models.City{})

	// make sure connection is available
	return db
}

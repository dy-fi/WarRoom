package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// MySql driver
	_ "github.com/go-sql-driver/mysql"
	// "github.com/dy-fi/war-room/models"
)

// Connect to MySql db
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "dev@tcp(:3333)/wrdb")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connected succesfully to db")
	}
	db.LogMode(true)
	//defer db.Close() // turn on for testing

	// Turn off in production, turn on in dev
	// db.AutoMigrate(&models.User{}, &models.Room{})

	// make sure connection is available
	return db
}

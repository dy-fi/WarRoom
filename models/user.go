package models

import (
	"github.com/jinzhu/gorm"

	// "github.com/dy-fi/war-room/db"
)

// var d = db.Connect()

// User model
type User struct {
	gorm.Model
	email	 	string	`gorm:"UNIQUE; NOT NULL;"`
	password	string
	rooms		[]Room	
}

// Verify a username password pair and return the result as a bool
// func Verify(username, password string) bool {
// 	return false
// }
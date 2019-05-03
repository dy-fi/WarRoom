package models

import (
	"github.com/jinzhu/gorm"

	"gopkg.in/validator.v2"
)

// User model
type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;"`
	username string `gorm:"unique; not null;" validate:"min:8"`
	email    string `gorm:"unique; not null;" validate:"regexp=\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z"`
	password string `gorm:"NOT NULL;" validate:"min:8"`
	// authorization signal
	auth	 int	
	// owned rooms
	rooms []Room
}

// Functions here and not in the repository because they 
// should be exposed only when the model is accessed

// Validate that a user struct is valid according to the tags
func (u *User) Validate() (bool, error) {
	if err := validator.Validate(u); err != nil {
		return false, err
	}
	return true, nil
}

// Verify a username password pair and return the result as a bool
// func Verify(username, password string) bool {
// 	return false
// }

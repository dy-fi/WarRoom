package models

import (
	"github.com/jinzhu/gorm"

	"gopkg.in/validator.v2"
)

// User model
type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;" json:"id"`
	Username string `gorm:"unique; not null;" json:"username" validate:"min:8"`
	Email    string `gorm:"unique; not null;" json:"email" validate:"regexp=\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z"`
	Password string `gorm:"NOT NULL;" validate:"min:8"`
	// authorization signal
	Auth	 int	
	// owned rooms
	MyCities []City	`gorm:"foreignkey:Mycities"`
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

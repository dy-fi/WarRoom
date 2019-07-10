package models

import (
	"github.com/jinzhu/gorm"
)

// City model
type City struct {
	gorm.Model
	// name of the room
	Name		string		`json:"Name"`
	// original author, will show in parent and child docs
	Author		string		`gorm:"not null;" json:"author"`
	// scrapping locations
	Places		[]Location	`gorm:"default:'', foreignkey:ID;" json:"places"`
}
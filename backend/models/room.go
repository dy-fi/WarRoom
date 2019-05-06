package models

import (
	"github.com/jinzhu/gorm"
)

// Room model
type Room struct {
	gorm.Model
	ID			uint 		`gorm:"primary_key;" json:"ID"`
	// name of the room
	Name		string		`json:"Name"`
	// owner of this specific Room
	Owner		uint		`gorm:"not null;" json:"owner"`
	// original author, will show in parent and child docs
	Author		uint		`gorm:"not null;" json:"author"`
	// scrapping locations
	Locations 	[]string	`json:"locations"`
	// custom html blob
	Blob		string		`json:"blob"`
}
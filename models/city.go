package models

import (
	"github.com/jinzhu/gorm"
)

// City model
type City struct {
	gorm.Model
	// name of the room
	Name		string		`json:"Name"`
	// owner of this specific Room
	// Owner		uint		`gorm:"not null;" json:"owner"`
	// original author, will show in parent and child docs
	// Author		uint		`gorm:"not null;" json:"author"`
	// scrapping locations
	Pages		[]Page		`json:"urls"`
	// custom html blob
	Blob		string		`json:"blob"`
}
package models

import (
	"github.com/jinzhu/gorm"
)

// Location is a webscrapper address
type Location struct {
	gorm.Model
	// Key ~ name
	Key		string		`json:"name"`
	// Address is the tag identifying the location
	Address string		`json:"address"`
}
package models

import (
	"github.com/jinzhu/gorm"
)

// Location is a webscrapper address
type Location struct {
	gorm.Model
	// Key ~ name
	Key string `json:"name"`
	// URL is the requested url
	URL string `json:"url"`
	// Address is the Xpath identifying the location
	Address string `json:"address"`
	// State is the last recorded values for this location (10 to start)
	State []string `json:"state"`
	// Style is the graph style
	Style string
}

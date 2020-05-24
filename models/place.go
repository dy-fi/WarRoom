package models

import "github.com/jinzhu/gorm"
// "github.com/lib/pq"

// Place is a webscrapper address
type Place struct {
	gorm.Model
	// Name is the location reference
	Name string `json:"name"`
	// URL is the requested url
	URL string `json:"url"`
	// Address is the Xpath identifying the location
	Address string `json:"address"`
	// Style is the graph style
	Style string `json:"style"`
}

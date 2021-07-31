package models

import "github.com/jinzhu/gorm"
// "github.com/lib/pq"

// Place is a webscrapper address
type Place struct {
	gorm.Model
	// Room is the room id reference
	Room uint
	// Name is the location reference
	Name string `json:"name"`
	// URL is the requested url
	URL string `json:"url"`
	// Address is the Xpath identifying the location
	Address string `json:"address"`
	// Style is the graph style
	Style string `json:"style"`
}

// Point is a returned single live value from a place
type Point struct {
	gorm.Model
	// Place_Name is the returned values place
	Place_Name string `json:"place_name"`
	// Value is the returned value itself
	Value string `json:"value"`
}


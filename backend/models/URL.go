package models

import (
	// "github.com/jinzhu/gorm"
)

// URL is all the locations of the URL in the 
type URL struct {
	// Glob is the starting URL
	Glob		string	`json:"glob"`
	// Locations is the identifying tag in the page
	Locations	[]Location	`json:"locations"`
}
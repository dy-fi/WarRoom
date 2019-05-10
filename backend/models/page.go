package models

// "github.com/jinzhu/gorm"

// Page is the URL all the locations of the URL
type Page struct {
	// Glob is the starting URL
	Glob string `json:"glob"`
	// Locations is the identifying tag in the page
	Locations []Location `json:"locations"`
}

package models

import (
	"github.com/jinzhu/gorm"
	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

// Room model
type Room struct {
	gorm.Model
	ID			uint 		`gorm:"PRIMARY_KEY;" json:"ID"`
	// stored by ID
	Owner		uint		`gorm:"NOT NULL;" json:"owner"`
	Author		uint		`gorm:"NOT NULL;" json:"author"`
	Urls 		[]string	`json:"urls"`
}
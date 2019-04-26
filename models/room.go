package models

import (
	"github.com/jinzhu/gorm"
	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

// Room model
type Room struct {
	gorm.Model
	ID		uint
	urls 	[]string
}
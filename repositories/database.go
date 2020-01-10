package repos

import (
	"fmt"

	// mysql import
	// _ "github.com/go-sql-driver/mysql"

	// "github.com/jinzhu/gorm/dialects/postgres"//
	// "github.com/jinzhu/gorm"

	"github.com/globalsign/mgo"
)

// Connect to MySql db
func Connect() *mgo.Database {
	session, err := mgo.Dial("mongodb://db:27017/wrdb")
	if err != nil {
		fmt.Println(err)
	}

	db := session.DB("wrdb")
	if err = session.Ping(); err != nil {
		fmt.Print("Couldn't connect to db")
	}
	mgo.SetDebug(true)

	return db
}

// DB is the database reference
var DB *mgo.Database = Connect()

package repos

import (

	// mysql import
	// _ "github.com/go-sql-driver/mysql"

	// gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"

	// "github.com/globalsign/mgo"
)

// Connect to mongo db
// func Connect() *mgo.Database {
// 	session, err := mgo.Dial("mongodb://db:27017/wrdb")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	db := session.DB("wrdb")
// 	if err = session.Ping(); err != nil {
// 		fmt.Print("Couldn't connect to db")
// 	}
// 	mgo.SetDebug(true)

// 	return db
// }

// // DB is the database reference
// var DB *mgo.Database = Connect()

// Connect returns the DB client object
func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=wrdb")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	return db
}

// DB object exposed to package
var DB = Connect()

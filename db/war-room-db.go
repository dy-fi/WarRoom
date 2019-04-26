package db

import (
	"fmt"

	"database/sql"
	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect to MySql db
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "dev@tcp(:3333)/wrdb")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connected succesfully to db")
	}
	
	//defer db.Close() // turn on for testing

	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println("db connection error")
		fmt.Println(err.Error())
	}
	return db
}

package handlers

import (
	// "net/http"
	"github.com/labstack/echo"

	// "github.com/dy-fi/war-room/models"
	"github.com/dy-fi/war-room/db"
	// "net/http"
)

// DB is the database reference
var DB = db.Connect()

// GetRooms handler - get all rooms associated with a user
// func GetRooms(c echo.Context) error {
// 	// get user id
// 	user := c.FormValue("user")
// 	// query for rooms
// 	rooms,err := DB.Query("select rooms from users WHERE ID == ?", user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return nothing
// 	// return c.Render(http.StatusOK, "rooms.html", rooms)
// 	return nil
// }

// GetRoom handler - get one room
func GetRoom(c echo.Context) error {
	return nil 
}

// PutRoom handler - update one room
func PutRoom(c echo.Context) error {
	return nil
}

// PostRoom handler - create one room
func PostRoom(c echo.Context) error {
	return nil
}

// DeleteRoom handler - delete one room
func DeleteRoom(c echo.Context) error {
	return nil 
}
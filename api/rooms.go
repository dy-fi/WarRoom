package api

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/dy-fi/war-room/models"
)

// GetRooms handler - get all rooms associated with a user
func GetRooms(c echo.Context) error {
	// get user id
	id := c.FormValue("user")
	// query and bind user
	user := models.User{}
	DB.First(&user, id)
	// query and bind rooms
	rooms := []models.Room{}
	DB.Where("owner = ?", id).Find(&rooms)

	return c.Render(http.StatusOK, "rooms", rooms)
}

// GetRoom handler - get one room by ID
func GetRoom(c echo.Context) error {
	// get id
	id := c.Param("id")
	// query and bind
	room := models.Room{}
	DB.Where("ID == ?", id).Find(&room)

	return c.Render(http.StatusOK, "room-show", room)
}

// PutRoom handler - update one room by ID
func PutRoom(c echo.Context) error {
	// get ID
	id := c.Param("id")
	// query and bind
	room := models.Room{}
	DB.Where("ID == ?", id).Find(&room)
	// Bind new values to room
	if err := c.Bind(room); err != nil {
		return err 
	}

	return c.Redirect(http.StatusAccepted, "/")
}

// PostRoom handler - create one room
func PostRoom(c echo.Context) error {
	r := models.Room{}
	
	if err := c.Bind(r); err != nil {
		return err
	}
	DB.Create(r)
	return c.Redirect(http.StatusCreated, "/")
}

// DeleteRoom handler - delete one room by ID
func DeleteRoom(c echo.Context) error {
	id := c.Param("id")
	r := models.Room{}
	DB.Where("ID == ?", id).Find(&r)
	DB.Delete(&r)

	return c.Redirect(http.StatusAccepted, "/")
}
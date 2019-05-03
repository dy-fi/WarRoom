package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/dy-fi/war-room/models"
	"github.com/dy-fi/war-room/repositories"
)

func GetAllRooms(c echo.Context) error {
	
}

// GetRooms handler - get all rooms associated with a user
func GetRooms(c echo.Context) error {
	// get user id
	id := c.Param("user")
	rooms, err := repos.GetRoomsByOwner(id)
	// handle errors
	if err != nil {
		log.Println(err)
		return c.Render(http.StatusOK, "rooms", "failure")
	}
	return c.Render(http.StatusOK, "rooms", rooms)
}

// NewRoomForm - render new room form
func NewRoomForm(c echo.Context) error {
	return c.Render(http.StatusOK, "partials/room-form", echo.Map{
		"user": nil,
	})
}

// GetRoom handler - get one room by ID
func GetRoom(c echo.Context) error {
	// get id
	id := c.Param("id")
	// get room
	room, err := repos.GetRoomByID(id)
	if err != nil {
		log.Println(err)
		return c.Render(http.StatusOK, "rooms", "failure")
	}

	return c.Render(http.StatusOK, "room-show", room)
}

// EditRoom handler - update one room by ID
func EditRoom(c echo.Context) error {
	// get values
	id := c.Param("id")
	name := c.FormValue("name")
	locs := c.FormValue("locations")
	blob := c.FormValue("Blob")

	// get room
	room, err := repos.GetRoomByID(id)
	if err != nil {
		log.Println(err)
		return c.Render(http.StatusOK, "rooms", "Couldn't find room")
	}
	// bind
	room.Name = name
	room.Locations = strings.Fields(locs)
	room.Blob = blob

	return c.Redirect(http.StatusAccepted, "/")
}

// CreateRoom handler - create one room
func CreateRoom(c echo.Context) error {
	r := models.Room{}

	if err := c.Bind(r); err != nil {
		return c.Render(http.StatusOK, "/rooms", "Couldn't create room")
	}
	room, err := repos.CreateRoom(r)
	if err != nil {

	}
	return c.Redirect(http.StatusCreated, "/rooms" + repos.GetID(room))
}

// DeleteRoom handler - delete one room by ID
func DeleteRoom(c echo.Context) error {
	id := c.Param("id")
	room,err := repos.GetRoomByID(id)
	if err != nil {
		return c.Redirect(http.StatusBadRequest, "/rooms/" + repos.GetID(room))
	}
	if repos.DeleteRoom(room) != nil {
		return c.Redirect(http.StatusInternalServerError, "/rooms/" + repos.GetID(room))
	}

	return c.Redirect(http.StatusAccepted, "/")
}

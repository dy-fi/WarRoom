package handlers

import (
	"log"
	"net/http"
	
	"github.com/labstack/echo"

	"github.com/dy-fi/war-room/models"
	repos "github.com/dy-fi/war-room/repositories"
)


// GetAllRooms - get all room documents in database
func GetAllRooms(c echo.Context) error {
	rooms, err := repos.GetAllRooms()

	if err != nil {
		log.Println(err)
		return c.Render(http.StatusInternalServerError, "error", err)
	}

	return c.JSON(http.StatusOK, rooms)
}

// GetRooms handler - get all rooms associated with a user
func GetRooms(c echo.Context) error {
	// get user id
	id := c.Param("user")
	rooms, err := repos.GetRoomsByOwner(id)
	// handle errors
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, "Error: Couldn't locate user document")
	}

	return c.JSON(http.StatusOK, rooms)
}

// GetRoom handler - get one room by ID
func GetRoom(c echo.Context) error {
	// get id
	id := c.Param("id")
	// get room
	room, err := repos.GetRoomByID(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, "Error: Couldn't get room")
	}

	return c.JSON(http.StatusOK, room)
}

// EditRoom handler - update one room by ID
func EditRoom(c echo.Context) error {
	// get ID param
	id := c.Param("id")
	// get room
	room, err := repos.GetRoomByID(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find room")
	}
	// bind
	if err := c.Bind(&room).Error; err != nil {

	}

	return c.JSON(http.StatusAccepted, room)
}

// CreateRoom handler - create one room
func CreateRoom(c echo.Context) error {
	r := models.Room{}
	// couldn't bind request data to model
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Couldn't bind data to room model")
	}
	// couldn't create room in database
	room, err := repos.CreateRoom(r)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't create room in database")
	}
	return c.JSON(http.StatusCreated, room)
}

// DeleteRoom handler - delete one room by ID
func DeleteRoom(c echo.Context) error {
	id := c.Param("id")
	room, err := repos.GetRoomByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find room to delete")
	}
	if repos.DeleteRoom(room) != nil {
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't delete room")
	}

	return c.JSON(http.StatusAccepted, "/")
}

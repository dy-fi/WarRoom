package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	// "github.com/dy-fi/war-room/models"
	repos "github.com/dy-fi/war-room/repositories"
)

// GetLocationByID retrieves a location given its ID in the request body
func GetLocationByID(c echo.Context) error {
	// get user id
	id, err := repos.StringToUint(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't resolve ID")
	}
	cities, err := repos.GetCityByID(id)
	// handle errors
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, "Error: Couldn't locate user document")
	}

	return c.JSON(http.StatusOK, cities)
}
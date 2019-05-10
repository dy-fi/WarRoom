package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/dy-fi/war-room/models"
	repos "github.com/dy-fi/war-room/repositories"
)

// Getcity handler - starts websocket connection and reads city
// func Getcity(c echo.Context) error {
// 	// get id
// 	id, err := repos.StringToUint(c.Param("id"))
// 	if err != nil {
// 		log.Printf("Couldnt resolve id: %v\n", err)
// 		return c.JSON(http.StatusBadRequest, "Error: Couldn't resolve city id")
// 	}
// 	// get city
// 	city, err := repos.GetCityByID(id)
// 	if err != nil {
// 		log.Println(err)
// 		return c.JSON(http.StatusOK, "Error: Couldn't get city")
// 	}
// 	// sockets

// }

// GetAllCities - get all city documents in database
func GetAllCities(c echo.Context) error {
	cities, err := repos.GetAllCities()

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cities)
}

// GetCities handler - get all cities associated with a user
func GetCities(c echo.Context) error {
	// get user id
	id, err := repos.StringToUint(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't resolve ID")
	}
	cities, err := repos.GetCitiesByOwner(id)
	// handle errors
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, "Error: Couldn't locate user document")
	}

	return c.JSON(http.StatusOK, cities)
}

// EditCity handler - update one city by ID
func EditCity(c echo.Context) error {
	// get ID param
	id, err := repos.StringToUint(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't resolve ID")
	}
	// get city
	city, err := repos.GetCityByID(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find city")
	}
	// bind
	if err := c.Bind(&city).Error; err != nil {

	}

	return c.JSON(http.StatusAccepted, city)
}

// CreateCity handler - create one city
func CreateCity(c echo.Context) error {
	r := models.City{}
	// couldn't bind request data to model
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Couldn't bind data to city model")
	}
	// couldn't create city in database
	city, err := repos.CreateCity(r)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't create city in database")
	}
	return c.JSON(http.StatusCreated, city)
}

// DeleteCity handler - delete one city by ID
func DeleteCity(c echo.Context) error {
	id, err := repos.StringToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error: couldn't resolve ID")
	}
	city, err := repos.GetCityByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find city to delete")
	}
	if repos.DeleteCity(city) != nil {
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't delete city")
	}

	return c.JSON(http.StatusAccepted, "/")
}

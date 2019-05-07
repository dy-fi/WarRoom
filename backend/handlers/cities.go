package handlers

import (
	"log"
	"net/http"
	
	"github.com/labstack/echo"

	"github.com/dy-fi/war-room/models"
	repos "github.com/dy-fi/war-room/repositories"
)


// GetAllCities - get all city documents in database
func GetAllCities(c echo.Context) error {
	cities, err := repos.GetAllcities()

	if err != nil {
		log.Println(err)
		return c.Render(http.StatusInternalServerError, "error", err)
	}

	return c.JSON(http.StatusOK, cities)
}

// Getcities handler - get all cities associated with a user
func Getcities(c echo.Context) error {
	// get user id
	id := c.Param("user")
	cities, err := repos.GetcitiesByOwner(id)
	// handle errors
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, "Error: Couldn't locate user document")
	}

	return c.JSON(http.StatusOK, cities)
}

// Getcity handler - get one city by ID
func Getcity(c echo.Context) error {
	// get id
	id := c.Param("id")
	// get city
	city, err := repos.GetcityByID(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, "Error: Couldn't get city")
	}

	return c.JSON(http.StatusOK, city)
}

// Editcity handler - update one city by ID
func Editcity(c echo.Context) error {
	// get ID param
	id := c.Param("id")
	// get city
	city, err := repos.GetcityByID(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find city")
	}
	// bind
	if err := c.Bind(&city).Error; err != nil {

	}

	return c.JSON(http.StatusAccepted, city)
}

// Createcity handler - create one city
func Createcity(c echo.Context) error {
	r := models.city{}
	// couldn't bind request data to model
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Couldn't bind data to city model")
	}
	// couldn't create city in database
	city, err := repos.Createcity(r)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't create city in database")
	}
	return c.JSON(http.StatusCreated, city)
}

// Deletecity handler - delete one city by ID
func Deletecity(c echo.Context) error {
	id := c.Param("id")
	city, err := repos.GetcityByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find city to delete")
	}
	if repos.Deletecity(city) != nil {
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't delete city")
	}

	return c.JSON(http.StatusAccepted, "/")
}

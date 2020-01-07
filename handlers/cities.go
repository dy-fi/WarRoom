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

// GetAllCities - get all cities in database
func GetAllCities(c echo.Context) error {
	cities := repos.GetAllCities()

	return c.Render(http.StatusOK, "rooms.html", cities)
}

// GetCityByID - get a single city by its ID
func GetCityByID(c echo.Context) error {
	id := c.Param("id")

	city, err := repos.GetCityByID(id)
	if err != nil {
		log.Println("Error: ")
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find that city")
	}
	return c.JSON(http.StatusOK, city)
}

// // GetCities handler - get all cities associated with a user
// func GetCitiesByUser(c echo.Context) error {
// 	// get user id
// 	id, err := repos.StringToUint(c.Param("id"))
// 	if err != nil {
// 		log.Println(err)
// 		return c.JSON(http.StatusBadRequest, "Error: Couldn't resolve ID")
// 	}
// 	cities, err := repos.GetCitiesByOwner(id)
// 	// handle errors
// 	if err != nil {
// 		log.Println(err)
// 		return c.JSON(http.StatusOK, "Error: Couldn't locate user document")
// 	}

// 	return c.JSON(http.StatusOK, cities)
// }

// EditCity handler - update one city by ID
func EditCity(c echo.Context) error {
	// get ID param
	id := c.Param("id")

	repos.UpdateCity(id, c.FormParams)

	return c.Render(http.StatusAccepted, "/target/"+id, nil)
}

// CreateCity handler - create one city
func CreateCity(c echo.Context) error {
	r := new(models.City)

	// Manually set values
	r = &models.City{
		Name: c.FormValue("name"),
		// Place: TODO
	}

	// couldn't create city in database
	city, err := repos.CreateCity(*r)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't create city in database")
	}
	return c.JSON(http.StatusCreated, city)
}

// DeleteCity handler - delete one city by ID
func DeleteCity(c echo.Context) error {
	city, err := repos.GetCityByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't find city to delete")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Couldn't find city to delete")
	}
	if repos.DeleteCity(city) != nil {
		return c.JSON(http.StatusInternalServerError, "Error: Couldn't delete city")
	}

	return c.JSON(http.StatusAccepted, "/")
}

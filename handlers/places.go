package handlers

// import (
// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo"

// 	// "github.com/dy-fi/war-room/models"
// 	repos "github.com/dy-fi/war-room/repositories"
// )

// // GetLocationByID retrieves a location given its ID in the request body
// func GetLocationByID(c echo.Context) error {
// 	// get user id
// 	cities, err := repos.GetCityByID(c.Param("id"))
// 	if err != nil {
// 		log.Println(err)
// 		return c.JSON(http.StatusBadRequest, "Error: Couldn't resolve ID")
// 	}

// 	return c.JSON(http.StatusOK, cities)
// }

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/dy-fi/war-room/db"
	"github.com/dy-fi/war-room/handlers"
)

func main() {
	// ==================== INIT ==================== //
	e := echo.New()
	e.Static("/", "public")

	// ==================== MIDDLEWARE ==================== //
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// // Authentication
	// e.Use(middleware.BasicAuth(func(username, password string, e echo.Context) (bool, error) {
	// 	if models.Verify(username, password) {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	// ==================== ROUTES ==================== //
	// rooms
	e.GET("/cities/all", handlers.GetAllCities)
	e.GET("/cities/ws/:id", handlers.GetCity)
	e.GET("/cities", handlers.GetCities)
	e.PUT("/cities/:id/edit", handlers.EditCity)
	e.POST("/cities", handlers.CreateCity)
	e.DELETE("/cities/:id", handlers.DeleteCity)
	// authentication

	// ==================== START ==================== //
	e.Logger.Fatal(e.Start(":8000"))
	db.Connect()
}

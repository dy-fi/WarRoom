package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
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
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
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
	e.GET("/rooms/:id", handlers.GetRoom)
	e.GET("/rooms", handlers.GetRooms)
	e.PUT("/rooms/:id", handlers.EditRoom)
	e.POST("/rooms", handlers.CreateRoom)
	e.DELETE("/room/:id", handlers.DeleteRoom)
	// authentication

	// ==================== START ==================== //
	e.Logger.Fatal(e.Start(":8000"))
}

package main

import (
	//"errors"
	"io"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/dy-fi/war-room/handlers"
)

// Renderer serves templates with the pongo engine
type Renderer struct {
	Debug bool
}

// Render is the Pongo2 renderer
func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	var ctx pongo2.Context
	if data != nil {
		var ok bool
		ctx, ok = data.(pongo2.Context)
		if !ok {
			ctx = nil 
		}
	}

	var t *pongo2.Template
	var err error
	if r.Debug {
		t, err = pongo2.FromFile(name)
	} else {
		t, err = pongo2.FromCache(name)
	}

	if err != nil {
		return err
	}

	return t.ExecuteWriter(ctx, w)
}

func main() {

	// ================== RENDERER ================== //
	renderer := Renderer{
		Debug: true,
	}
	// ==================== INIT ==================== //
	e := echo.New()
	e.Renderer = renderer
	e.Static("/static", "static")
	// ================= MIDDLEWARE ================= //
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
	e.GET("/", handlers.Index)
	// e.GET("/rooms", handlers.GetCitiesByUser)
	// e.GET("/rooms/:id", handlers.GetCityByID)
	e.GET("/cities/ws/:id", handlers.GetCity)
	e.PUT("/cities/:id/edit", handlers.EditCity)
	e.GET("/room/new", handlers.NewCityForm)
	e.POST("/room/new", handlers.CreateCity)
	e.DELETE("/cities/:id", handlers.DeleteCity)
	// authentication

	// ==================== START ==================== //
	e.Logger.Fatal(e.Start(":8000"))
}

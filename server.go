package main

import (
	"html/template"
	"io"
	"os"

    "github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/dy-fi/war-room/handlers"
)

// Template reference struct
type Template struct {
	templates *template.Template
}

// Render Templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	// echo instance
	e := echo.New()
	e.Static("/", "public/views")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// template renderer
	e.Renderer = t

	// Routes
	// index
	e.GET("/", handlers.GetIndex)
	// rooms
	// e.GET("/rooms", handlers.GetRooms)

	// Start
	e.Logger.Fatal(e.Start(":8000"))
}

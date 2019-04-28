package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/dy-fi/war-room/db"

)

// DB is the database reference - this is shared to the
var DB = db.Connect()

// GetIndex index handler
func GetIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

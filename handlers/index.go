package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

// GetIndex index handler
func GetIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "index")
}

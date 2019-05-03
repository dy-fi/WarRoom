package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetIndex index handler
func GetIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

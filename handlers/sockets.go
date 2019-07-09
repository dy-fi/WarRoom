package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"

	repos "github.com/dy-fi/war-room/repositories"
)

var upgrader = websocket.Upgrader{}

// GetCity handler - starts websocket connection and reads city by id param
func GetCity(c echo.Context) error {
	// get id
	id, err := repos.StringToUint(c.Param("id"))
	if err != nil {
		log.Printf("Couldnt resolve id: %v\n", err)
		return c.JSON(http.StatusBadRequest, "Error: Couldn't resolve city id")
	}
	// get city
	city, err := repos.GetCityByID(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, "Error: Couldn't get city")
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteJSON(city)
		if err != nil {
			c.Logger().Error(err)
		}
		err2 := ws.WriteJSON(repos.ScrapeCity(&city))
		if err != nil {
			c.Logger().Error(err2)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		if string(msg) == "close" {
			ws.Close()
		}
		log.Printf("%s\n", msg)
	}
}

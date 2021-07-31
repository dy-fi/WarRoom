package repos

import (
	"github.com/dy-fi/war-room/models"
)

// Locate takes a place and returns a data representation as a string
func Locate(place models.Place) models.Point {
	  return models.Point{Place_Name: place.Name, Value: "1"}
}

// StreamPlace calls the service and sends data to a receiving channel
func StreamPlace(place models.Place, channel chan<- models.Point) {
	channel <- Locate(place)
}
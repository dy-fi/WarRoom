package repos

import (
	"github.com/dy-fi/war-room/models"
)

// Locate takes a place and returns a data representation as a string
func Locate(place models.Place) Point {
	
}

// StreamPlace calls the service and sends data to a receiving channel
func StreamPlace(place models.Place, channel chan<- Point) {
	channel <- Locate(place)
}
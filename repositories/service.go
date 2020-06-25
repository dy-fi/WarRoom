package repos

import (
	"github.com/dy-fi/war-room/models"
)

// CallService reaches out the arbitrary service for place data
func CallService(place models.Place) {
	
}

// StreamPlace calls the service and sends data to a receiving channel
func StreamPlace(place model.Place, channel chan<- Point) {
	channel <- callservice(place)
}
package repos

import (
	"errors"

	"github.com/dy-fi/war-room/models"
	
)


// TODO

// CreatePlace creates a place in the database 
func CreatePlace(place models.Place) (models.Place, error) {
	// check that row does not exist
	if DB.NewRecord(place) {
		DB.Create(&place)
	} else {
		return place, errors.New("Place already exists")
	}

	return place, nil
}

// GetPlaceID returns a string formatted id of the given city
func GetPlaceID(place models.Place) string {
	return string(place.ID)
}

// GetPlacesByRoomID returns a list of all places in a room given the room id
func GetPlacesByRoomID(rid uint) []models.Place {
	var result []models.Place
	DB.Where("Room = ?", rid).Find(&result)
	return result
}

// UpdatePlace updates a place
func UpdatePlace(place *models.Place, updates map[string]interface{}) {
	DB.Model(&place).Updates(updates)
}

// DeletePlace removes a place
func DeletePlace(place models.Place) error {
	// hard delete
	if err := DB.Delete(&place); err != nil {
		return err.Error
	}

	return nil 
}




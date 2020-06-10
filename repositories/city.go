package repos

import (
	"errors"

	"github.com/dy-fi/war-room/models"
)

// GetCityID returns an Id given a city model
func GetCityID(city models.City) string {
	return string(city.ID)
}

// GetAllCities returns a list of every city document in the database
func GetAllCities() []models.City {
	var result []models.City
	DB.Find(&result)
	return result
}

// UpdateCity updates a city
func UpdateCity(city *models.City, updates map[string]interface{}) {
	DB.Model(&city).Updates(updates)
}

// // GetCitiesByOwner gets all cities owned by a user indexed by user ID - has many
// func GetCitiesByOwner(id uint) ([]models.City, error) {
// 	cities := []models.City{}
// 	user := models.User{}
// 	D.First(&user, id)
// 	if err := D.Model(&user).Related(&cities).Error; err != nil {
// 		return nil, err
// 	}
// 	return cities, nil
// }

// GetCityByID gets a city document indexed by ID
func GetCityByID(id string) (models.City, error) {
	city := models.City{}
	if err := DB.First(&city, id); err != nil {
		return city, errors.New("Couldn't find city")
	}

	return city, nil
}

// CreateCity makes a new city
func CreateCity(city models.City) (models.City, error) {
	if DB.NewRecord(city) {
		DB.Create(&city)
	} else {
		return city, errors.New("City already exists")
	}

	return city, nil
}

// DeleteCity removes a city
func DeleteCity(city models.City) error {
	// hard delete
	if err := DB.Delete(&city); err != nil {
		return err.Error
	}

	return nil 
}


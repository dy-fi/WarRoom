package repos

import (
	"errors"

	"github.com/dy-fi/war-room/models"
)

DB.AutoMigrate(&models.City)

// GetCityID returns an Id given a city model
func GetCityID(city models.City) string {
	return string(city.Id)
}

// GetAllCities returns a list of every city document in the database
func GetAllCities() []models.City {
	result := []models.City{}
	if err := DB.Find(&city); err != nil {
		panic(err.Error())
	}
	return result
}

// UpdateCity updates a city
func UpdateCity(city *models.City) {
	DB.Save(&city)
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
	if err := cities.FindId(id).One(&city); err != nil {
		return city, errors.New("Couldn't find city")
	}

	return city, nil
}

// CreateCity makes a new city
func CreateCity(city models.City) (models.City, error) {
	if DB.NewRecord(city) {
		DB.Create(&city)
	}
}

// DeleteCity removes a city
func DeleteCity(city models.City) error {
	// hard delete
	db.Delete(&city)
}


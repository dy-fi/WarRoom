package repos

import (
	"log"
	"strconv"

	"github.com/dy-fi/war-room/models"
)

// StringToUint helper for id string to uint
func StringToUint(s string) (uint, error) {
	intID, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Couldnt resolve id: %v\n", err)
		return uint(0), err
	}
	return uint(intID), nil
}

// GetAllCities returns a list of every city document in the database
func GetAllCities() ([]models.City, error) {
	cities := []models.City{}
	if err := D.Find(&cities).Error; err != nil {
		return cities, err
	}
	return cities, nil
}

// GetID of a city
func GetID(city models.City) string {
	return strconv.Itoa(int(city.ID))
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
func GetCityByID(id uint) (models.City, error) {
	city := models.City{}
	if err := D.First(&models.City{}, "ID = ?", id).Error; err != nil {
		return city, err
	}
	return city, nil
}

// CreateCity makes a new city
func CreateCity(city models.City) (models.City, error) {
	if err := D.Create(&city).Error; err != nil {
		return city, err
	}
	return city, nil
}

// DeleteCity removes a city
func DeleteCity(city models.City) error {
	if err := D.Delete(&city).Error; err != nil {
		return err
	}
	return nil
}

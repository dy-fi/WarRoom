package repos

import (
	"errors"

	"github.com/dy-fi/war-room/models"
)

var cities = DB.C("cities")

// GetCityID returns an Id given a city model
func GetCityID(city models.City) string {
	return string(city.Id)
}

// GetAllCities returns a list of every city document in the database
func GetAllCities() []models.City {
	result := []models.City{}
	if err := cities.Find(nil).All(&result); err != nil {
		panic(err.Error())
	}
	return result
}

// UpdateCity updates a city
func UpdateCity(id string, data interface{}) error {
	if err := cities.UpdateId(id, data); err != nil {
		return err
	}
	return nil
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
	if err := cities.Insert(&city); err != nil {
		return city, err
	}
	return city, nil
}

// DeleteCity removes a city
func DeleteCity(city models.City) error {
	if err := cities.RemoveId(city.Id); err != nil {
		return err
	}
	return nil
}


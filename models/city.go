package models

import (
	"github.com/flosch/pongo2"
	"github.com/jinzhu/gorm"
)

// City model
type City struct {
	gorm.Model
	// name of the room
	Name string `json:"name"`
	// original author, will show in parent and child docs
	Author string `json:"author"`
}

// ToPongoContext converts a city struct into a pongo readable format
func (c City) ToPongoContext() pongo2.Context {
	data := pongo2.Context{
		"name": c.Name,
		"author": c.Author,
	}
	return data

}

// ListToPongo takes a list of cities and returns them in pongo readable format
func ListToPongo(cities []City) []pongo2.Context {
	result := []pongo2.Context{}

	if cities != nil {
		for _,v := range cities {
			result = append(result, v.ToPongoContext()) 
		}
		return result 
	} else {
		return result 
	}
	
}
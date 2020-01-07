package models

import (
	"github.com/globalsign/mgo/bson"

)

// City model
type City struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// name of the room
	Name string `json:"name"`
	// original author, will show in parent and child docs
	Author string `json:"author"`
	// scrapping locations
	Places []bson.ObjectId `json:"places"`
	// output type
	Output bool `json:"output"`
}

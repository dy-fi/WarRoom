package repos

import (
	"github.com/dy-fi/war-room/database"
)

// DB is the database client object
var DB = database.Connect()

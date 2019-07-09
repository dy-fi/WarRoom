package repos

import (
	"strconv"

	"github.com/dy-fi/war-room/models"
)

// GetUserByID returns a user document, indexed by the given ID
func GetUserByID(id string) (models.User, error) {
	uid, _ := strconv.ParseUint(id, 10, 32)
	user := models.User{}
	if err := D.First(&user, uid).Error; err != nil {
		return user, err
	}
	return user, nil
}
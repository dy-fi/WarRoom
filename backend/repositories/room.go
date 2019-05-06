package repos

import (
	"strconv"

	"github.com/dy-fi/war-room/models"
)

// GetAllRooms returns a list of every room document in the database
func GetAllRooms() ([]models.Room, error) {
	rooms := []models.Room{}
	if err := D.Find(&rooms).Error; err != nil {
		return rooms, err
	}
	return rooms, nil
}

// GetID of a room
func GetID(room models.Room) string {
	return strconv.Itoa(int(room.ID))
}

// GetRoomsByOwner gets all rooms owned by a user indexed by user ID
func GetRoomsByOwner(oid string) ([]models.Room, error) {
	rooms := []models.Room{}
	if err := D.Where("owner = ?", oid).Find(&rooms).Error; err != nil {
		return rooms, err
	}
	return rooms, nil
}

// GetRoomByID gets a room document indexed by ID
func GetRoomByID(id string) (models.Room, error) {
	uid, _ := strconv.ParseUint(id, 10, 32)
	room := models.Room{}
	if err := D.First(&room, uid).Error; err != nil {
		return room, err
	}
	return room, nil
}

// CreateRoom makes a new room
func CreateRoom(room models.Room) (models.Room, error) {
	if err := D.Create(room).Error; err != nil {
		return room, err
	}
	return room, nil
}

// DeleteRoom removes a room 
func DeleteRoom(room models.Room) error {
	if err := D.Delete(&room).Error; err != nil {
		return err
	}
	return nil
}
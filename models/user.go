package models

// User model
type User struct {
	ID			string
	username 	string
	rooms		[]Room
}
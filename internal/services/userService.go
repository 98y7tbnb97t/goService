package services

import (
	"echoServer/internal/repositories"
	"echoServer/models"
)

type User = models.User

func GetAllUsers() ([]User, error) {
	return repositories.GetAllUsers()
}

// CreateUser inserts a new user.
func CreateUser(user *User) error {
	return repositories.CreateUser(user)
}

// GetUserByID retrieves a user by ID.
func GetUserByID(id int) (*User, error) {
	return repositories.GetUserByID(id)
}

// DeleteUser deletes a user by ID.
func DeleteUser(id int) error {
	return repositories.DeleteUser(id)
}

// UpdateUser updates a user with new values.
func UpdateUser(id int, updates *User) error {
	return repositories.UpdateUser(id, updates)
}

func GetTasksForUser(userID uint) ([]models.Task, error) {
	return repositories.GetTasksForUser(userID)
}

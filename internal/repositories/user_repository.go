package repositories

import (
	"echoServer/db"
	"echoServer/models"
)

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
}

// CreateUser inserts a new user into the database.
func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

// GetUserByID retrieves a user by its ID.
func GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser removes a user from the database.
func DeleteUser(id int) error {
	return db.DB.Delete(&models.User{}, id).Error
}

// UpdateUser updates an existing user using the provided updates.
func UpdateUser(id int, updates *models.User) error {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return err
	}
	return db.DB.Model(&user).Updates(updates).Error
}

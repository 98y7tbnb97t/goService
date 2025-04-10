package repository

import (
	"echoServer/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(offset, limit int) ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id uint, user models.User) (models.User, error)
	DeleteUser(id uint) error
	// CountUsers returns the total number of users.
	CountUsers() (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers(offset, limit int) ([]models.User, error) {
	var users []models.User
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(id uint, user models.User) (models.User, error) {
	var existing models.User
	if err := r.db.First(&existing, id).Error; err != nil {
		return models.User{}, err
	}

	if err := r.db.Model(&existing).Updates(user).Error; err != nil {
		return models.User{}, err
	}
	return existing, nil
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// CountUsers returns the total number of users.
func (r *userRepository) CountUsers() (int64, error) {
	var count int64
	if err := r.db.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

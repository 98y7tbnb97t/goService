package service

import (
	"echoServer/internal/userService/repository"
	"echoServer/models"

	"golang.org/x/crypto/bcrypt"
)

// UserService defines an interface for user-related operations
type UserService interface {
	GetUsers(page, limit int) ([]models.User, int64, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id uint, user models.User) (models.User, error)
	DeleteUser(id uint) error
}

// userServiceImpl implements the UserService interface
type userServiceImpl struct {
	repo repository.UserRepository
}

// NewUserService returns a new instance of UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetUsers(page, limit int) ([]models.User, int64, error) {
	offset := (page - 1) * limit

	// Use the CountUsers method from the repository
	total, err := s.repo.CountUsers()
	if err != nil {
		return nil, 0, err
	}

	users, err := s.repo.GetUsers(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userServiceImpl) CreateUser(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)

	return s.repo.CreateUser(user)
}

func (s *userServiceImpl) UpdateUser(id uint, user models.User) (models.User, error) {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, err
		}
		user.Password = string(hashedPassword)
	}
	return s.repo.UpdateUser(id, user)
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

// DummyUserService is a mock implementation of UserService for testing
type DummyUserService struct{}

// NewDummyUserService returns a new instance of DummyUserService
func NewDummyUserService() UserService {
	return &DummyUserService{}
}

// GetUsers returns dummy users data
func (s *DummyUserService) GetUsers(page, limit int) ([]models.User, int64, error) {
	dummyUsers := []models.User{
		{ID: 1, Email: "user1@example.com"},
		{ID: 2, Email: "user2@example.com"},
	}
	return dummyUsers, int64(len(dummyUsers)), nil
}

// CreateUser returns a dummy created user
func (s *DummyUserService) CreateUser(user models.User) (models.User, error) {
	user.ID = 999 // Dummy ID
	return user, nil
}

// UpdateUser returns a dummy updated user
func (s *DummyUserService) UpdateUser(id uint, user models.User) (models.User, error) {
	user.ID = id
	return user, nil
}

// DeleteUser is a dummy implementation that always succeeds
func (s *DummyUserService) DeleteUser(id uint) error {
	return nil
}

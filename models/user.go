package models

import (
	"fmt"
	"time"
	"unicode/utf8"

	"gorm.io/gorm"
)

// User представляет модель пользователя в базе данных
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null;size:255" json:"email" validate:"required,email"`
	Password  string         `gorm:"not null;size:255" json:"password" validate:"required,min=8"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Tasks     []Task         `gorm:"foreignKey:UserID" json:"tasks,omitempty"` // Связь с задачами
}

// TableName возвращает имя таблицы для модели User.
func (User) TableName() string {
	return "users"
}

// Validate проверяет корректность данных пользователя
func (u *User) Validate() error {
	if utf8.RuneCountInString(u.Email) == 0 {
		return fmt.Errorf("email is required")
	}

	if utf8.RuneCountInString(u.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	return nil
}

// BeforeSave хук, который вызывается перед сохранением пользователя
func (u *User) BeforeSave(tx *gorm.DB) error {
	return u.Validate()
}

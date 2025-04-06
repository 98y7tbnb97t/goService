package models

import "time"

type Task struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Task      string    `json:"task"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

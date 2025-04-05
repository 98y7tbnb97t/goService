package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Task{})
}

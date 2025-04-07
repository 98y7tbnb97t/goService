package main

import (
	"echoServer/db"
	"echoServer/internal/handlers"
	"echoServer/internal/middleware"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()

	if err := db.DB.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	e := echo.New()

	e.Use(middleware.TimestampMiddleware)

	handlers.RegisterTaskHandlers(e)

	log.Println("Server starting on port :8882")
	if err := e.Start(":8882"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

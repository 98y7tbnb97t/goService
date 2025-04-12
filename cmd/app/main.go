package main

import (
	"echoServer/db"
	"echoServer/internal/handlers/tasksHandlers" // Import tasks handlers package
	"echoServer/internal/handlers/userHandlers"  // Import user handlers package
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

	// Register tasks related endpoints from tasksHandlers (renamed from handlers)
	tasksHandlers.RegisterTaskHandlers(e)
	// Register user related endpoints from the new userHandlers
	userHandlers.RegisterUserHandlers(e)

	log.Println("Server starting on port :8882")
	if err := e.Start(":8882"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"log"

	"echoServer/db"
	"echoServer/internal/handlers"
	"echoServer/internal/middleware"
	"echoServer/internal/repositories"
	"echoServer/internal/services"
	userRepoPkg "echoServer/internal/userService/repository"
	userServicePkg "echoServer/internal/userService/service"
	"echoServer/models"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()

	// Auto migration for all models using models package
	if err := db.DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	e := echo.New()
	e.Use(middleware.TimestampMiddleware)

	// Initialize task components
	taskRepo := repositories.NewGormRepository(db.DB)
	taskSer := services.NewService(taskRepo)
	handlers.RegisterTaskHandlers(e, taskSer)

	// Initialize user components
	userRepo := userRepoPkg.NewUserRepository(db.DB)
	userSer := userServicePkg.NewUserService(userRepo)
	handlers.RegisterUserHandlers(e, userSer)

	log.Println("Server starting on port :8882")
	if err := e.Start(":8882"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"echoServer/db"
	"echoServer/internal/handlers"
	"echoServer/internal/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()

	e := echo.New()

	e.Use(middleware.TimestampMiddleware)

	handlers.RegisterTaskHandlers(e)

	e.Start(":8882")
}

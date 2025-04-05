package main

import (
	"echoServer/db"
	"echoServer/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()

	e := echo.New()

	handlers.RegisterTaskHandlers(e)

	e.Start(":8882")
}

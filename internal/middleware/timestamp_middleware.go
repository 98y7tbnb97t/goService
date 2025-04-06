package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"echoServer/models"
)

func TimestampMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" || c.Request().Method == "PUT" || c.Request().Method == "PATCH" {
			task := new(models.Task)
			if err := c.Bind(task); err == nil {
				now := time.Now()
				if c.Request().Method == "POST" {
					task.CreatedAt = now
				}
				task.UpdatedAt = now
			}
		}
		return next(c)
	}
}
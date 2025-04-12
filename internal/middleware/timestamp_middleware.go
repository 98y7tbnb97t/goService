package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

func TimestampMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" || c.Request().Method == "PUT" || c.Request().Method == "PATCH" || c.Request().Method == "DELETE" {
			now := time.Now()
			c.Set("current_time", now)
		}
		return next(c)
	}
}

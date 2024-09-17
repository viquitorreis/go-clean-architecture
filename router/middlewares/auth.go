package middlewares

import (
	"log"
	"time"

	"github.com/labstack/echo"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()

		// Call the next handler
		err := next(c)

		// Log the request details
		method := c.Request().Method
		path := c.Request().URL.Path
		duration := time.Since(startTime)

		log.Printf("Method: %s, Path: %s, Duration: %v", method, path, duration)

		return err
	}
}

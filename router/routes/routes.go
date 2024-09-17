package routes

import (
	"cleanarch/router/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitECHO() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	NirRouteManager(e)
	e.Logger.Fatal(e.Start(":6969"))
}

func NirRouteManager(e *echo.Echo) {
	g := e.Group("/nir")
	g.GET("/results", handlers.GetNIRResults)
	g.GET("/result/:id", handlers.GetNIRResultsByID)
}

package main

import (
	"learn-go-echo/src/controllers"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // Middleware
	e.Use(middleware.Logger()) // Logger
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api/v1", serverHeader)
	api.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	api.POST("/price", controllers.GrabPrice) // Price endpoint

	// Server
	e.Logger.Fatal(e.Start(":8000"))
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v1.0")
		return next(c)
	}
}

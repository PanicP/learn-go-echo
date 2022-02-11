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
	// Root route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.POST("/price", controllers.GrabPrice) // Price endpoint
	// Server
	e.Logger.Fatal(e.Start(":8000"))
}

// The "godef" command is not available. Run "go install -v github.com/rogpeppe/godef@latest" to install.

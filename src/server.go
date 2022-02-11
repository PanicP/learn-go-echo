package main
import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "net/http"
  )
func main() {
 
 e := echo.New() 
 // Middleware
 e.Use(middleware.Logger())
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
// Run Server
e.Logger.Fatal(e.Start(":8000"))
}
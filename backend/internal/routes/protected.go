package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitProtectedRoutes(e *echo.Echo) {
	// Initialize the public routes here
	g := e.Group("/api")
	g.Use(echojwt.JWT([]byte("secret-key")))

}

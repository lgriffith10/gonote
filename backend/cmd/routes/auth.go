package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/controllers"
)

func InitAuthRoutes(e *echo.Echo) {
	authController := controllers.AuthController{}

	g := e.Group("/auth")
	g.POST("/login", authController.Login)
}

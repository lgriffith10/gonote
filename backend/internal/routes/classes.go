package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/controllers"
)

func InitClassesRoute(e *echo.Echo) {
	classController := controllers.NewClassController()

	g := e.Group("/api/classes")

	g.POST("/", classController.CreateClass)
}

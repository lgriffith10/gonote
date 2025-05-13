package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lgriffith10/gonote/cmd/routes"
	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/validators"
)

func main() {
	database.InitDatabase()

	e := echo.New()
	e.Validator = &validators.CustomValidator{Validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitAuthRoutes(e)
	routes.InitProtectedRoutes(e)

	e.Logger.Fatal(e.Start(":4000"))
}

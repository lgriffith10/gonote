package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/routes"
	"github.com/lgriffith10/gonote/internal/translations"
	"github.com/lgriffith10/gonote/internal/validators"
)

func main() {
	e := SetupServer(false)

	e.Logger.Fatal(e.Start(":4000"))
}

func SetupServer(isTest bool) *echo.Echo {
	database.InitDatabase(isTest)
	e := echo.New()
	e.Validator = &validators.CustomValidator{Validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
	}))

	translations.LoadTranslations()

	routes.InitAuthRoutes(e)
	routes.InitProtectedRoutes(e)

	return e
}

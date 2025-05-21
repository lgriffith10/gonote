package test

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/validators"
)

func SetupTestServer() *echo.Echo {
	e := echo.New()
	e.Validator = &validators.CustomValidator{Validator: validator.New()}

	return e
}

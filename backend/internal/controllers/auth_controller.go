package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

func (a *AuthController) Login(c echo.Context) error {
	result, err := a.authService.Login(c)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (a *AuthController) Register(c echo.Context) error {
	err := a.authService.RegisterUser(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

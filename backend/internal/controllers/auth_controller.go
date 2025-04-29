package controllers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	auth_jwt "github.com/lgriffith10/gonote/internal/auth"
)

type AuthController struct{}

type LoginRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (a *AuthController) Login(c echo.Context) error {
	r := new(LoginRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if r.Name != "jon" || r.Password != "password" {
		return echo.ErrUnauthorized
	}

	claims := &auth_jwt.JwtCustomClaims{
		Name:  "Jon Snow",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

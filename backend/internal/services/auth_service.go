package services

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	auth_jwt "github.com/lgriffith10/gonote/internal/auth"
	"github.com/lgriffith10/gonote/internal/dtos"
	"github.com/lgriffith10/gonote/internal/models"
	"github.com/lgriffith10/gonote/internal/repositories"
	"github.com/lgriffith10/gonote/internal/utils"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
}

func NewAuthService() *AuthService {
	authRepository := repositories.NewAuthRepository()
	return &AuthService{
		authRepository: authRepository,
	}
}

func (a *AuthService) Login(c echo.Context) (*dtos.LoginResponse, error) {
	r := new(dtos.LoginRequest)
	err := utils.ParseRequest(c, r)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	areCredentialsCorrect := a.authRepository.AreCredentialsCorrect(r.Email, r.Password)

	if !areCredentialsCorrect {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := &auth_jwt.JwtCustomClaims{
		Name:  r.Email,
		Admin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret-key"))

	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{Token: t}, nil
}

func (a *AuthService) RegisterUser(c echo.Context) error {
	r := new(dtos.RegisterRequest)
	err := utils.ParseRequest(c, r)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user := models.User{
		Email:     r.Email,
		Password:  r.Password,
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
	}

	err = a.authRepository.RegisterUser(user)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}

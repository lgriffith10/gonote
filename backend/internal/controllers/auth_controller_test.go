package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/dtos"
	"github.com/lgriffith10/gonote/internal/dtos/builders"
	"github.com/lgriffith10/gonote/internal/models"
	"github.com/lgriffith10/gonote/internal/repositories"
	"github.com/lgriffith10/gonote/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	database.InitDatabase(true)
	e := utils.SetupTestServer(t)

	user := models.User{
		Email:     "test@test.com",
		Password:  "password",
		Firstname: "Luciano",
		Lastname:  "Griffith",
	}

	userRepository := repositories.NewAuthRepository()

	userRepository.RegisterUser(user)

	t.Cleanup(func() {
		database.DB().Where(&user).Delete(&user)
	})

	t.Run("should return 200 status", func(t *testing.T) {
		loginRequest := builders.NewLoginRequestBuilder().WithEmail(user.Email).WithPassword(user.Password).Get()
		rec, err := sendAndTreatRequest(t, e, loginRequest)

		if err != nil {
			t.Fatalf("expected nil, got err: %v", err)
		}

		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		data, _ := utils.ReadResponseBody(t, rec.Result(), &dtos.LoginResponse{})

		assert.NotNil(t, data.Token)
	})

	t.Run("should return 401 status with invalid credentials", func(t *testing.T) {
		loginRequest := builders.NewLoginRequestBuilder().WithEmail("bep@bep.com").Get()
		rec, err := sendAndTreatRequest(t, e, loginRequest)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
		data, _ := utils.ReadResponseError(t, rec.Result())

		assert.Equal(t, "Invalid credentials", data)
	})

	t.Run("should return 422 status with invalid request", func(t *testing.T) {
		loginRequest := builders.NewLoginRequestBuilder().WithEmail("").WithPassword("").Get()

		rec, err := sendAndTreatRequest(t, e, loginRequest)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func sendAndTreatRequest(t testing.TB, e *echo.Echo, request *dtos.LoginRequest) (*httptest.ResponseRecorder, error) {
	t.Helper()
	body, err := utils.ParseRequestToBytes(request)

	if err != nil {
		t.Fatalf("failed to parse request to bytes: %v", err)
	}

	req := httptest.NewRequest(echo.POST, "/auth/login", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controller := NewAuthController()
	err = controller.Login(c)

	if err != nil {
		e.HTTPErrorHandler(err, c)
	}

	return rec, err
}

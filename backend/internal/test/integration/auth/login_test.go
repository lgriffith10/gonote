package auth_integration_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/controllers"
	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/dtos"
	"github.com/lgriffith10/gonote/internal/models"
	"github.com/lgriffith10/gonote/internal/repositories"
	"github.com/lgriffith10/gonote/internal/test/builders"
	"github.com/lgriffith10/gonote/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	database.InitDatabase(true)
	e := utils.SetupTestServer(t)

	user := models.User{
		Id:        uuid.New(),
		Email:     "test@test.com",
		Password:  "password",
		Firstname: "Luciano",
		Lastname:  "Griffith",
	}

	userRepository := repositories.NewAuthRepository()

	userRepository.RegisterUser(user)

	t.Cleanup(func() {
		database.DB().Delete(&user)
	})

	t.Run("should return 200 status", func(t *testing.T) {
		loginRequest := builders.NewLoginRequestBuilder().WithEmail(user.Email).WithPassword(user.Password).Get()
		rec, err := sendAndTreatLoginRequest(t, e, loginRequest)

		if err != nil {
			t.Fatalf("expected nil, got err: %v", err)
		}

		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		data, _ := utils.ReadResponseBody(t, rec.Result(), &dtos.LoginResponse{})

		assert.NotNil(t, data.Token)
	})

	t.Run("should return 401 status with invalid credentials", func(t *testing.T) {
		loginRequest := builders.NewLoginRequestBuilder().WithEmail("bep@bep.com").Get()
		rec, err := sendAndTreatLoginRequest(t, e, loginRequest)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
		data, _ := utils.ReadResponseError(t, rec.Result())

		assert.Equal(t, "Invalid credentials", data)
	})

	t.Run("should return 422 status with invalid request", func(t *testing.T) {
		loginRequest := builders.NewLoginRequestBuilder().WithEmail("").WithPassword("").Get()

		rec, err := sendAndTreatLoginRequest(t, e, loginRequest)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func sendAndTreatLoginRequest(t testing.TB, e *echo.Echo, request *dtos.LoginRequest) (*httptest.ResponseRecorder, error) {
	c, rec := utils.ParseAndSendRequest(t, e, request, "/auth/login")

	controller := controllers.NewAuthController()
	err := controller.Login(c)

	if err != nil {
		e.HTTPErrorHandler(err, c)
	}

	return rec, err
}

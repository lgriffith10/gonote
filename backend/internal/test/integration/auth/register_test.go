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

func TestRegister(t *testing.T) {
	database.InitDatabase(true)
	_context := utils.SetupTestServer(t)

	_userRepository := repositories.NewAuthRepository()

	_user := &models.User{
		Id:        uuid.New(),
		Email:     "bep@bep.com",
		Password:  "password",
		Firstname: "Luciano",
		Lastname:  "Griffith",
	}

	_userRepository.RegisterUser(_user)

	t.Cleanup(func() {
		database.DB().Delete(_user)
	})

	t.Run("Register with valid data should work", func(t *testing.T) {
		request := builders.NewRegisterRequestBuilder().Get()
		rec, err := sendAndTreatRegisterRequest(t, _context, request)

		if err != nil {
			t.Fatalf("expected nil, got err: %v", err)
		}

		var user models.User
		database.DB().Debug().Where("email = ?", request.Email).First(&user)

		assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)
		assert.Equal(t, request.Email, user.Email)
		assert.Equal(t, request.Firstname, user.Firstname)
		assert.Equal(t, request.Lastname, user.Lastname)
		assert.NotNil(t, user.Id)

		t.Cleanup(func() {
			database.DB().Debug().Delete(&user)
		})
	})

	t.Run("Register with existing email should return 409", func(t *testing.T) {
		request := builders.NewRegisterRequestBuilder().WithEmail(_user.Email).Get()
		rec, err := sendAndTreatRegisterRequest(t, _context, request)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		assert.Equal(t, http.StatusConflict, rec.Result().StatusCode)
	})

	t.Run("Register with invalid data should return 400", func(t *testing.T) {
		request := builders.NewRegisterRequestBuilder().WithEmail("invalid-email").Get()
		rec, err := sendAndTreatRegisterRequest(t, _context, request)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func sendAndTreatRegisterRequest(t testing.TB, e *echo.Echo, request *dtos.RegisterRequest) (*httptest.ResponseRecorder, error) {
	c, rec := utils.ParseAndSendRequest(t, e, request, "/auth/login")

	controller := controllers.NewAuthController()
	err := controller.Register(c)

	if err != nil {
		e.HTTPErrorHandler(err, c)
	}

	return rec, err
}

package classes_integration_test

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
	"github.com/lgriffith10/gonote/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateClass(t *testing.T) {
	_context := utils.SetupTestServer(t)
	_userRepository := repositories.NewAuthRepository()

	_manager := &models.User{
		Id:        uuid.New(),
		Email:     "test@test.com",
		Password:  "password",
		Firstname: "Test",
		Lastname:  "User",
	}

	_userRepository.RegisterUser(_manager)

	t.Run("createClass should work with valid data", func(t *testing.T) {
		request := &dtos.CreateClassRequest{
			Name: "Test class",
		}
		rec, err := sendAndTreatCreateClassRequest(t, _context, request)

		if err != nil {
			t.Fatalf("expected nil, got err: %v", err)
		}

		res, _ := utils.ReadResponseBody(t, rec.Result(), &dtos.CreateClassResponse{})

		class := &models.Class{}
		database.DB().Debug().First(class, res.Id)

		assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)
		assert.NotNil(t, res.Id)

		assert.Equal(t, class.Name, request.Name)
	})
}

func sendAndTreatCreateClassRequest(t testing.TB, e *echo.Echo, request *dtos.CreateClassRequest) (*httptest.ResponseRecorder, error) {
	c, rec := utils.ParseAndSendRequest(t, e, request, "/api/classes/")

	controller := controllers.NewClassController()
	err := controller.CreateClass(c)

	if err != nil {
		e.HTTPErrorHandler(err, c)
	}

	return rec, err
}

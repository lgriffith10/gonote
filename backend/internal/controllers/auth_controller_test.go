package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	database.InitDatabase(true)

	t.Run("should return 200 status", func(t *testing.T) {
		e := test.SetupTestServer()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		controller := NewAuthController()
		err := controller.Login(c)

		println(database.DB())

		if err == nil {
			t.Error("expected error, got nil")
		} else {
			e.HTTPErrorHandler(err, c)
		}

		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

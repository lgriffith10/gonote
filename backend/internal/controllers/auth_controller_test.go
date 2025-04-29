package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("should return 200 status", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/auth/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		controller := AuthController{}
		controller.Login(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

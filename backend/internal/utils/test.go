package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/translations"
	"github.com/lgriffith10/gonote/internal/validators"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SetupTestServer(t testing.TB) *echo.Echo {
	t.Helper()

	database.InitDatabase(true)
	e := echo.New()
	e.Validator = &validators.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Recover())

	translations.LoadTranslations()

	return e
}

func ReadResponseBody[T any](t testing.TB, res *http.Response, str T) (T, error) {
	t.Helper()

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	json.Unmarshal(data, &str)

	return str, err
}

func ReadResponseError(t testing.TB, res *http.Response) (string, error) {
	t.Helper()

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	var str ErrorResponse
	json.Unmarshal(data, &str)

	return str.Message, err
}

package utils

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func ParseRequest[T any](c echo.Context, r *T) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}

func ParseRequestToBytes(request any) ([]byte, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return body, nil
}

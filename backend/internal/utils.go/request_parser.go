package utils

import (
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

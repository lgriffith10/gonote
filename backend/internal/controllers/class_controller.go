package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/services"
)

type classController struct {
	classService *services.ClassService
}

func NewClassController() *classController {
	return &classController{
		classService: services.NewClassService(),
	}
}

func (cl *classController) CreateClass(c echo.Context) error {
	res, err := cl.classService.CreateClass(c)

	if err != nil {
		return err
	}

	return c.JSON(201, res)
}

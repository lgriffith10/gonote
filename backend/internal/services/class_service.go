package services

import (
	"github.com/labstack/echo/v4"
	"github.com/lgriffith10/gonote/internal/dtos"
	"github.com/lgriffith10/gonote/internal/repositories"
)

type ClassService struct {
	repository *repositories.ClassRepository
}

func NewClassService() *ClassService {
	return &ClassService{
		repository: repositories.NewClassRepository(),
	}
}

func (cs *ClassService) CreateClass(c echo.Context) (*dtos.CreateClassResponse, error) {
	res, err := cs.repository.CreateClass()

	return &dtos.CreateClassResponse{
		Id: res,
	}, err
}

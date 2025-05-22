package dtos

import "github.com/google/uuid"

type CreateClassRequest struct {
	Name      string      `json:"name" validate:"required,min=1,max=100"`
	Students  []uuid.UUID `json:"students" validate:"required"`
	ManagerId uuid.UUID   `json:"manager_id" validate:"required,uuid"`
}

type CreateClassResponse struct {
	Id uuid.UUID `json:"id"`
}

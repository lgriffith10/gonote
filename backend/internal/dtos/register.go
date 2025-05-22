package dtos

type RegisterRequest struct {
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}

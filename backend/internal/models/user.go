package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Classes   []Class   `gorm:"foreignKey:ManagerId;references:Id"`
}

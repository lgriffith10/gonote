package models

import "github.com/google/uuid"

type Class struct {
	Id        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Students  []User    `gorm:"many2many:class_students"`
	ManagerId uuid.UUID
	Manager   User `gorm:"foreignKey:ManagerId"`
}

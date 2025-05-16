package models

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

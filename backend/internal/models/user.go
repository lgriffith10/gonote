package models

type User struct {
	Id        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Classes   []Class `gorm:"foreignKey:ManagerId;references:Id"`
}

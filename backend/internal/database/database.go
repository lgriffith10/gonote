package database

import (
	"fmt"

	"github.com/lgriffith10/gonote/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func InitDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=gonote"
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully and migrated")
}

func DB() *gorm.DB {
	return database
}

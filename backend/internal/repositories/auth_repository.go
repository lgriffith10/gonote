package repositories

import (
	"errors"

	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/models"
	"github.com/lgriffith10/gonote/internal/translations"
	"github.com/lgriffith10/gonote/internal/utils"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		DB: database.DB(),
	}
}

func (a *AuthRepository) RegisterUser(user *models.User) error {
	isEmailAlreadyTaken := a.DB.Model(&models.User{}).Where("email = ?", user.Email).First(&user).RowsAffected == 1

	if isEmailAlreadyTaken {
		msg := translations.GetTranslation("emailAlreadyExists", nil)

		return errors.New(msg)
	}

	password, err := utils.HashPassword(user.Password)

	if err != nil {
		msg := translations.GetTranslation("registerUserError", nil)

		return errors.New(msg)
	}

	user.Password = password

	if err := a.DB.Create(&user).Error; err != nil {
		msg := translations.GetTranslation("registerUserError", nil)

		return errors.New(msg)
	}

	return nil
}

func (a *AuthRepository) AreCredentialsCorrect(email, password string) bool {
	var hash string
	result := a.DB.Model(&models.User{}).Where("email = ?", email).Select("password").First(&hash)

	if result.RowsAffected == 0 {
		return false
	}

	if isPasswordCorrect := utils.CheckPasswordHash(password, hash); !isPasswordCorrect {
		return false
	}

	return true
}

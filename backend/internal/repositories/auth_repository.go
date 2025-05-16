package repositories

import (
	"errors"

	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/models"
	"github.com/lgriffith10/gonote/internal/translations"
	"github.com/lgriffith10/gonote/internal/utils.go"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		db: database.DB(),
	}
}

func (a *AuthRepository) RegisterUser(user models.User) error {
	isEmailAlreadyTaken := a.db.Model(&models.User{}).Where("email = ?", user.Email).First(&user).RowsAffected == 1

	if isEmailAlreadyTaken {
		msg, _ := translations.GetTranslation("emailAlreadyExists", nil)

		return errors.New(msg)
	}

	if err := a.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (a *AuthRepository) AreCredentialsCorrect(email, password string) bool {
	var hash string
	result := a.db.Model(&models.User{}).Where("email = ?", email).Select("password").First(&hash)

	if result.RowsAffected == 0 {
		return false
	}

	if isPasswordCorrect := utils.CheckPasswordHash(password, hash); !isPasswordCorrect {
		return false
	}

	return true
}

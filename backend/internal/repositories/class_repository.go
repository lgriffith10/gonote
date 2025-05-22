package repositories

import (
	"github.com/lgriffith10/gonote/internal/database"
	"github.com/lgriffith10/gonote/internal/models"
	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository() *ClassRepository {
	return &ClassRepository{
		db: database.DB(),
	}
}

func (r *ClassRepository) CreateClass(c *models.Class) error {
	if err := r.db.Create(c).Error; err != nil {
		return err
	}

	return nil
}

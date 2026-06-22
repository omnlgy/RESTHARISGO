package repository

import (
	"github.com/omnlgy/RESTHARISGO/internal/models"
	"gorm.io/gorm"
)

type positionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *positionRepository {
	return &positionRepository{
		db: db,
	}
}

func (r *positionRepository) Create(position *models.Position) (models.Position, error) {
	return *position, r.db.Create(position).Error
}

func (r *positionRepository) GetAll() ([]models.Position, error) {
	var positions []models.Position
	return positions, r.db.Find(&positions).Error
}

func (r *positionRepository) Update(position *models.Position) (models.Position, error) {
	return *position, r.db.Save(position).Error
}

func (r *positionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Position{}, id).Error
}

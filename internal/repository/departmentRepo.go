package repository

import (
	"errors"

	"github.com/omnlgy/RESTHARISGO/internal/models"
	"gorm.io/gorm"
)

// ErrNotFound is returned when a record is not found.
var ErrNotFound = errors.New("record not found")

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{
		db: db,
	}
}

func (r *departmentRepository) GetAll() ([]models.Department, error) {
	var departments []models.Department
	return departments, r.db.Find(&departments).Error
}

func (r *departmentRepository) GetByID(id uint) (models.Department, error) {
	var department models.Department
	err := r.db.First(&department, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Department{}, ErrNotFound
	}
	return department, err
}

func (r *departmentRepository) Create(department *models.Department) (models.Department, error) {
	return *department, r.db.Create(department).Error
}

func (r *departmentRepository) Update(department *models.Department) (models.Department, error) {
	result := r.db.Save(department)

	if result.RowsAffected == 0 {
		return models.Department{}, ErrNotFound
	}

	return *department, result.Error
}

func (r *departmentRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Department{}, id)

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return result.Error
}

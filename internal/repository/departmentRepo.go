package repository

import (
	"github.com/omnlgy/RESTHARISGO/internal/models"
	"gorm.io/gorm"
)

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

func (r *departmentRepository) Create(department *models.Department) (models.Department, error) {
	return *department, r.db.Create(department).Error
}

func (r *departmentRepository) Update(department *models.Department) (models.Department, error) {
	return *department, r.db.Save(department).Error
}

func (r *departmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}

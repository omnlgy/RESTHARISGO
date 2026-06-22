package domain

import "github.com/omnlgy/RESTHARISGO/internal/models"

type DepartmentRepository interface {
	Create(department *models.Department) (models.Department, error)
	GetAll() ([]models.Department, error)
	GetByID(id uint) (models.Department, error)
	Update(department *models.Department) (models.Department, error)
	Delete(id uint) error
}

type PositionRepository interface {
	Create(position *models.Position) (models.Position, error)
	GetAll() ([]models.Position, error)
	GetByID(id uint) (models.Position, error)
	Update(position *models.Position) (models.Position, error)
	Delete(id uint) error
}

type EmployeeRepository interface {
	Create(employee *models.Employee) (models.Employee, error)
	GetAll() ([]models.Employee, error)
	GetByID(id uint) (models.Employee, error)
	Update(employee *models.Employee) (models.Employee, error)
	Delete(id uint) error
}

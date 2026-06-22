package domain

type DepartmentRepository interface {
	Create(department *models.Department) (models.Department, error)
	GetAll() ([]models.Department, error)
	Update(department *models.Department) (models.Department, error)
	Delete(id uint) error
}
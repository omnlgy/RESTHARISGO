package service

import (
	"github.com/omnlgy/RESTHARISGO/internal/domain"
	"github.com/omnlgy/RESTHARISGO/internal/models"
	"github.com/omnlgy/RESTHARISGO/internal/repository"
)

type DepartmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo domain.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repo: repo,
	}
}

func (s *DepartmentService) GetDepartments() ([]models.Department, error) {
	return s.repo.GetAll()
}

func (s *DepartmentService) CreateDepartment(department *models.Department) (models.Department, error) {
	createdDepartment, err := s.repo.Create(department)

	if err != nil {
		return models.Department{}, err
	}

	return createdDepartment, nil
}

func (s *DepartmentService) UpdateDepartment(department *models.Department) (models.Department, error) {
	updatedDepartment, err := s.repo.Update(department)

	if err != nil {
		return models.Department{}, err
	}

	return updatedDepartment, nil
}

func (s *DepartmentService) DeleteDepartment(id uint) error {
	return s.repo.Delete(id)
}

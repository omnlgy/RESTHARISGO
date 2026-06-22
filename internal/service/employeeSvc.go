package service

import (
	"github.com/omnlgy/RESTHARISGO/internal/domain"
	"github.com/omnlgy/RESTHARISGO/internal/models"
)

type EmployeeService struct {
	repo    domain.EmployeeRepository
	deptSvc *DepartmentService
	posSvc  *PositionService
}

func NewEmployeeService(repo domain.EmployeeRepository, deptSvc *DepartmentService, posSvc *PositionService) *EmployeeService {
	return &EmployeeService{
		repo:    repo,
		deptSvc: deptSvc,
		posSvc:  posSvc,
	}
}

func (s *EmployeeService) Add(employee *models.Employee) (models.Employee, error) {
	_, err := s.deptSvc.GetDepartmentByID(employee.DepartmentID)
	if err != nil {
		return models.Employee{}, err
	}

	_, err = s.posSvc.GetPositionByID(employee.PositionID)
	if err != nil {
		return models.Employee{}, err
	}

	return s.repo.Create(employee)
}

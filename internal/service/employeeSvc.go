package service

import (
	"errors"
	"fmt"

	"github.com/omnlgy/RESTHARISGO/internal/domain"
	"github.com/omnlgy/RESTHARISGO/internal/models"
)

type EmployeeService struct {
	repo    domain.EmployeeRepository
	deptSvc domain.DepartmentService
	posSvc  domain.PositionService
}

func NewEmployeeService(repo domain.EmployeeRepository, deptSvc domain.DepartmentService, posSvc domain.PositionService) *EmployeeService {
	return &EmployeeService{
		repo:    repo,
		deptSvc: deptSvc,
		posSvc:  posSvc,
	}
}

func (s *EmployeeService) Add(employee *models.Employee) (models.Employee, error) {
	dept, err := s.deptSvc.GetDepartmentByID(employee.DepartmentID)
	if err != nil {
		return models.Employee{}, err
	}

	pos, err := s.posSvc.GetPositionByID(employee.PositionID)
	if err != nil {
		return models.Employee{}, err
	}

	employee.Department = dept
	employee.Position = pos

	return s.repo.Create(employee)
}

func (s *EmployeeService) GetEmployees(filter models.FilterEmployee) ([]models.Employee, error) {
	return s.repo.GetAll(filter)
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
	return s.repo.Delete(id)
}

func (s *EmployeeService) Update(employee *models.Employee) (models.Employee, error) {
	dept, err := s.deptSvc.GetDepartmentByID(employee.DepartmentID)
	if err != nil {
		fmt.Println(err)
		return models.Employee{}, errors.New("Department not found")
	}

	pos, err := s.posSvc.GetPositionByID(employee.PositionID)
	if err != nil {
		fmt.Println(err)
		return models.Employee{}, errors.New("Position not found")
	}

	employee.Department = dept
	employee.Position = pos

	return s.repo.Update(employee)
}

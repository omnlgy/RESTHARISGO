package service

import (
	"github.com/omnlgy/RESTHARISGO/internal/domain"
	"github.com/omnlgy/RESTHARISGO/internal/models"
)

type PositionService struct {
	repo domain.PositionRepository
}

func NewPositionService(repo domain.PositionRepository) *PositionService {
	return &PositionService{
		repo: repo,
	}
}

func (s *PositionService) GetPositions() ([]models.Position, error) {
	return s.repo.GetAll()
}

func (s *PositionService) CreatePosition(position *models.Position) (models.Position, error) {
	createdDepartment, err := s.repo.Create(position)

	if err != nil {
		return models.Position{}, err
	}

	return createdDepartment, nil
}

func (s *PositionService) UpdatePosition(position *models.Position) (models.Position, error) {
	updatedDepartment, err := s.repo.Update(position)

	if err != nil {
		return models.Position{}, err
	}

	return updatedDepartment, nil
}

func (s *PositionService) DeletePosition(id uint) error {
	return s.repo.Delete(id)
}

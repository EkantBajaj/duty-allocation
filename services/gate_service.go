package services

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/repositories"
)

type GateService interface {
	CreateGateEntry(gateEntry *models.GateEntry) error
	GetTotalGateCount() (count int, err error)
	GetGateCountByGate() ([]models.GateCountWithName, error)
	GetGates() ([]models.Gate, error)
}

type gateService struct {
	gateRepo repositories.GateRepository
}

func NewGateService(gateRepo repositories.GateRepository) GateService {
	return &gateService{
		gateRepo: gateRepo,
	}
}

func (s *gateService) CreateGateEntry(gateEntry *models.GateEntry) error {
	err := s.gateRepo.CreateGateEntry(gateEntry)
	if err != nil {
		return err
	}
	return nil
}

func (s *gateService) GetTotalGateCount() (count int, err error) {
	count, err = s.gateRepo.GetTotalGateCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *gateService) GetGateCountByGate() ([]models.GateCountWithName, error) {
	count, err := s.gateRepo.GetGateCountByGate()
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (s *gateService) GetGates() ([]models.Gate, error) {
	gates, err := s.gateRepo.GetGates()
	if err != nil {
		return nil, err
	}
	return gates, nil
}
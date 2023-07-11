package services

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/repositories"
)

type SpotService interface {
	CreateSpot(spot *models.Spot) error
	GetSpot(u uint) (*models.Spot, error)
	GetAllSpots() ([]models.Spot, error)
}

type spotService struct {
	spotRepository *repositories.SpotRepository
}

func NewSpotService(spotRepository *repositories.SpotRepository) *spotService {
	return &spotService{spotRepository: spotRepository}
}

func (ss *spotService) CreateSpot(spot *models.Spot) error {
	return ss.spotRepository.Create(spot)
}

func (ss *spotService) GetSpot(id uint) (*models.Spot, error) {
	return ss.spotRepository.FindByID(id)
}

func (ss *spotService) GetAllSpots() ([]models.Spot, error) {
	return ss.spotRepository.GetAllSpots()
}

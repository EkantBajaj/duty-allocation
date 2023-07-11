package services

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/repositories"
)

type SpotUserService interface {
	CreateSpotUser(spotUser *models.SpotUser) error
	GetActiveSpotUserCount() ([]models.ActiveSpotUserCount, error)
	GetActiveUsersBySpotID(id uint) ([]models.ActiveUsersBySpotID, error)
	GetActiveSpotUserByID(userID uint) (*models.SpotUser, error)
	UpdateUser(spotUser *models.SpotUser) error
}

type spotUserService struct {
	spotUserRepository *repositories.SpotUserRepository
}

func NewSpotUserService(spotUserRepository *repositories.SpotUserRepository) *spotUserService {
	return &spotUserService{
		spotUserRepository: spotUserRepository,
	}
}

func (sus *spotUserService) CreateSpotUser(spotUser *models.SpotUser) error {
	return sus.spotUserRepository.CreateSpotUser(spotUser)
}

func (sus *spotUserService) GetActiveSpotUserCount() ([]models.ActiveSpotUserCount, error) {
	return sus.spotUserRepository.GetActiveSpotUserCount()
}

func (sus *spotUserService) GetActiveUsersBySpotID(id uint) ([]models.ActiveUsersBySpotID, error) {
	return sus.spotUserRepository.GetActiveUsersBySpotID(id)
}

func (sus *spotUserService) GetActiveSpotUserByID(userID uint) (*models.SpotUser, error) {
	return sus.spotUserRepository.GetActiveSpotUserByID(userID)
}

func (sus *spotUserService) UpdateUser(spotUser *models.SpotUser) error {
	return sus.spotUserRepository.UpdateSpotUser(spotUser)
}

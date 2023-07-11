package repositories

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"gorm.io/gorm"
)

type SpotRepository struct {
	db *gorm.DB
}

func NewSpotRepository(db *gorm.DB) *SpotRepository {
	return &SpotRepository{db: db}
}

func (sr *SpotRepository) Create(spot *models.Spot) error {
	return sr.db.Create(spot).Error
}

func (sr *SpotRepository) GetSpotById(id uint) (*models.Spot, error) {
	var spot models.Spot
	err := sr.db.First(&spot, id).Error
	return &spot, err
}

func (sr *SpotRepository) FindByID(id uint) (*models.Spot, error) {
	var spot models.Spot
	if err := sr.db.First(&spot, id).Error; err != nil {
		return nil, err
	}
	return &spot, nil
}

func (sr *SpotRepository) GetAllSpots() ([]models.Spot, error) {
	var spots []models.Spot
	err := sr.db.Find(&spots).Error
	if err != nil {
		return nil, err
	}
	return spots, nil
}

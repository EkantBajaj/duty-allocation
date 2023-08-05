package repositories

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"gorm.io/gorm"
)

type SpotUserRepository struct {
	db *gorm.DB
}

func NewSpotUserRepository(db *gorm.DB) *SpotUserRepository {
	return &SpotUserRepository{db: db}
}

func (sur *SpotUserRepository) CreateSpotUser(spotUser *models.SpotUser) error {
	return sur.db.Create(spotUser).Error
}

func (sur *SpotUserRepository) FindByID(id uint) (*models.SpotUser, error) {
	var spotUser models.SpotUser
	if err := sur.db.First(&spotUser, id).Error; err != nil {
		return nil, err
	}
	return &spotUser, nil
}
func (sur *SpotUserRepository) GetActiveSpotUserCount() ([]models.ActiveSpotUserCount, error) {
	var results []models.ActiveSpotUserCount

	err := sur.db.Model(&models.Spot{}).
		Select("spots.*, COUNT(spot_users.id) as user_count").
		Joins("LEFT JOIN spot_users ON spots.id = spot_users.spot_id").
		Where("spot_users.active = ?", true).
		Group("spots.id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (sur *SpotUserRepository) GetActiveUsersBySpotID(spotID uint) ([]models.ActiveUsersBySpotID, error) {
	var users []models.ActiveUsersBySpotID

	err := sur.db.
		Model(&models.User{}).
		Select("users.id, name, gender, badge_id, initiated,spot_users.in_time, spot_users.out_time, users.mobile_number,users.emergency_number").
		Joins("JOIN spot_users ON spot_users.user_id = users.id").
		Where("spot_users.spot_id = ? AND spot_users.active = ?", spotID, true).
		Scan(&users).Error
	if err != nil {
		return nil, err
	}

	for i := range users {
		users[i].InTimeString = users[i].InTime.Format("15:04:05")   // Customize the time format as needed
		users[i].OutTimeString = users[i].OutTime.Format("15:04:05") // Customize the time format as needed
	}

	return users, nil
}

// GetActiveSpotUserByID retrieves an active spot user by ID
func (sur *SpotUserRepository) GetActiveSpotUserByID(userId uint) (*models.SpotUser, error) {
	var spotUser models.SpotUser
	err := sur.db.Where("user_id = ? AND active = ?", userId, true).Preload("Spot").First(&spotUser).Error
	if err != nil {
		return nil, err
	}
	return &spotUser, nil
}

// UpdateSpotUser updates the spot user in the database
func (sur *SpotUserRepository) UpdateSpotUser(spotUser *models.SpotUser) error {
	err := sur.db.Save(spotUser).Error
	if err != nil {
		return err
	}
	return nil
}

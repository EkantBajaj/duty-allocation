package repositories

import (
	"errors"
	"github.com/ekantbajaj/duty-allocation/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Roles").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByMobileNumber(mobileNumber string) (*models.User, error) {
	var user models.User
	err := r.db.Where("mobile_number = ?", mobileNumber).Preload("Roles").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	err := r.db.Save(user).Error
	return err
}

func (r *UserRepository) DeleteUser(id uint) error {
	err := r.db.Delete(&models.User{}, id).Error
	return err
}

func (r *UserRepository) GetUserByBadgeID(badgeID string) (*models.User, error) {
	var user models.User
	result := r.db.Where("badge_id = ?", badgeID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

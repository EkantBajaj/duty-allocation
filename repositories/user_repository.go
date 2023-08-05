package repositories

import (
	"errors"
	"fmt"
	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/token"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db         *gorm.DB
	tokenMaker token.Maker
}

func NewUserRepository(db *gorm.DB, token token.Maker) *UserRepository {
	return &UserRepository{
		db:         db,
		tokenMaker: token,
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
			return nil, fmt.Errorf("user not found with badge id")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) CreateToken(badgeId string) (string, error) {
	duration, _ := time.ParseDuration(viper.GetString("token.expiration_duration"))
	accessToken, err := r.tokenMaker.CreateToken(badgeId, duration)
	return accessToken, err
}

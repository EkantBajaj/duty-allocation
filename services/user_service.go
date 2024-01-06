package services

import (
	"errors"

	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/repositories"
	"github.com/ekantbajaj/duty-allocation/util"
)

type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	GetUserByMobileNumber(mobileNumber string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetUserByBadgeId(badgeId string) (*models.User, error)
	CreateToken(badgeId string) (string, error)
}

type userService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByMobileNumber(mobileNumber string) (*models.User, error) {
	user, err := s.userRepository.GetUserByMobileNumber(mobileNumber)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateUser(user *models.User) error {
	var err error
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	// check if badge id is empty then make it first four digit of mobile number followed by last four digit of aadhar number
	if user.BadgeID == "" {
		user.BadgeID = user.MobileNumber[0:4] + user.AadharNumber[len(user.AadharNumber)-4:]
	}
	err = s.userRepository.CreateUser(user)
	return err
}

func (s *userService) UpdateUser(user *models.User) error {
	err := s.userRepository.UpdateUser(user)
	return err
}

func (s *userService) DeleteUser(id uint) error {
	err := s.userRepository.DeleteUser(id)
	return err
}

func (s *userService) GetUserByBadgeId(badgeId string) (*models.User, error) {
	user, err := s.userRepository.GetUserByBadgeID(badgeId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateToken(badgeId string) (string, error) {
	token, err := s.userRepository.CreateToken(badgeId)
	if err != nil {
		return "", err
	}
	return token, nil

}

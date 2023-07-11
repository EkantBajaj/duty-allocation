package services

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/repositories"
)

type RoleService interface {
	GetRoleByID(id uint) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (s *roleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.roleRepo.GetRoleByID(id)
}

func (s *roleService) GetRoleByName(name string) (*models.Role, error) {
	return s.roleRepo.GetRoleByName(name)
}

func (s *roleService) CreateRole(role *models.Role) error {
	return s.roleRepo.CreateRole(role)
}

func (s *roleService) UpdateRole(role *models.Role) error {
	return s.roleRepo.UpdateRole(role)
}

func (s *roleService) DeleteRole(id uint) error {
	return s.roleRepo.DeleteRole(id)
}

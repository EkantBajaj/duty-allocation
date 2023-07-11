package repositories

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoleByID(id uint) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) CreateRole(role *models.Role) error {
	err := r.db.Create(role).Error
	return err
}

func (r *roleRepository) UpdateRole(role *models.Role) error {
	err := r.db.Save(role).Error
	return err
}

func (r *roleRepository) DeleteRole(id uint) error {
	err := r.db.Delete(&models.Role{}, id).Error
	return err
}

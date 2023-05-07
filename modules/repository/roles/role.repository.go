package roles

import (
	"fmt"

	// er "github.com/berrylradianh/go-jewelry/modules/entity/roles"
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (roleRepo *Repository) GetAllRoles() (*[]eu.Role, error) {
	var roles []eu.Role
	if err := roleRepo.DB.Preload("Users", "deleted_at IS NULL").Find(&roles).Error; err != nil {
		return nil, err
	}

	return &roles, nil
}

func (roleRepo *Repository) GetRoleById(id int) (*eu.Role, error) {
	var role eu.Role
	if err := roleRepo.DB.Preload("Users", "deleted_at IS NULL").First(&role, id).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (roleRepo *Repository) CreateRole(role *eu.Role) error {
	if err := roleRepo.DB.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

func (roleRepo *Repository) UpdateRole(id int, role *eu.Role) error {
	result := roleRepo.DB.Model(&role).Where("id = ?", id).Omit("UpdatedAt").Updates(&role)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (roleRepo *Repository) DeleteRole(id int) error {
	if err := roleRepo.DB.Delete(&eu.Role{}, id).Error; err != nil {
		return err
	}

	return nil
}

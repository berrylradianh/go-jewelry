package roles

import (
	"fmt"

	er "github.com/berrylradianh/go-jewelry/modules/entity/roles"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (roleRepo *Repository) GetAllRoles() (*[]er.Role, error) {
	var roles []er.Role
	if err := roleRepo.DB.Preload("User", "deleted_at IS NULL").Find(&roles).Error; err != nil {
		return nil, err
	}

	return &roles, nil
}

func (roleRepo *Repository) GetRoleById(id int) (*er.Role, error) {
	var role er.Role
	if err := roleRepo.DB.Preload("User", "deleted_at IS NULL").First(&role, id).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (roleRepo *Repository) CreateRole(role *er.Role) error {
	if err := roleRepo.DB.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

func (roleRepo *Repository) UpdateRole(id int, role *er.Role) error {
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
	if err := roleRepo.DB.Delete(&er.Role{}, id).Error; err != nil {
		return err
	}

	return nil
}

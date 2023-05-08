package roles

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (roleRepo *Repository) GetAllRoles() (*[]e.Role, error) {
	var roles []e.Role
	if err := roleRepo.DB.Preload("Users", "deleted_at IS NULL").Find(&roles).Error; err != nil {
		return nil, err
	}

	return &roles, nil
}

func (roleRepo *Repository) GetRoleById(id int) (*e.Role, error) {
	var role e.Role
	if err := roleRepo.DB.Preload("Users", "deleted_at IS NULL").First(&role, id).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (roleRepo *Repository) CreateRole(role *e.Role) error {
	if err := roleRepo.DB.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

func (roleRepo *Repository) UpdateRole(id int, role *e.Role) error {
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
	if err := roleRepo.DB.Delete(&e.Role{}, id).Error; err != nil {
		return err
	}

	return nil
}

package users

import (
	"fmt"

	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
)

func (userRepo *Repository) GetAllUsers() (*[]eu.User, error) {
	var users []eu.User
	if err := userRepo.DB.Preload("User_detail", "deleted_at IS NULL").Preload("Role", "deleted_at IS NULL").Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (userRepo *Repository) GetUserById(id int) (*eu.User, error) {
	var user eu.User
	if err := userRepo.DB.Preload("User_detail", "deleted_at IS NULL").Preload("Role", "deleted_at IS NULL").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *Repository) CreateUser(user *eu.User) error {
	if err := userRepo.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *Repository) UpdateUser(id int, user *eu.User) error {
	result := userRepo.DB.Model(&user).Where("id = ?", id).Omit("UpdatedAt").Updates(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product category with id %d not found", id)
	}

	return nil
}

func (userRepo *Repository) DeleteUser(id int) error {
	if err := userRepo.DB.Where("user_id = ?", id).Delete(&eu.UserDetail{}).Error; err != nil {
		return err
	}

	if err := userRepo.DB.Delete(&eu.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

package users

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (userRepo *Repository) GetAllUsers() (*[]e.User, error) {
	var users []e.User
	if err := userRepo.DB.Preload("User_detail", "deleted_at IS NULL").Preload("Role", "deleted_at IS NULL").Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (userRepo *Repository) GetUserById(id int) (*e.User, error) {
	var user e.User
	if err := userRepo.DB.Preload("User_detail", "deleted_at IS NULL").Preload("Role", "deleted_at IS NULL").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *Repository) CreateUser(user *e.User) error {
	if err := userRepo.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *Repository) UpdateUser(id int, user *e.User) error {
	result := userRepo.DB.Model(&user).Where("id = ?", id).Omit("UpdatedAt").Updates(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (userRepo *Repository) DeleteUser(id int) error {
	if err := userRepo.DB.Where("user_id = ?", id).Delete(&e.UserDetail{}).Error; err != nil {
		return err
	}

	if err := userRepo.DB.Delete(&e.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

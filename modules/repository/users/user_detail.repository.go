package users

import (
	"fmt"

	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
)

func (userDetailRepo *Repository) GetAllUserDetails() (*[]eu.UserDetail, error) {
	var userDetails []eu.UserDetail
	if err := userDetailRepo.DB.Preload("User", "deleted_at IS NULL").Find(&userDetails).Error; err != nil {
		return nil, err
	}

	return &userDetails, nil
}

func (userDetailRepo *Repository) GetUserDetailById(id int) (*eu.UserDetail, error) {
	var UserDetail eu.UserDetail
	if err := userDetailRepo.DB.Preload("User", "deleted_at IS NULL").First(&UserDetail, id).Error; err != nil {
		return nil, err
	}

	return &UserDetail, nil
}

func (userDetailRepo *Repository) CreateUserDetail(UserDetail *eu.UserDetail) error {
	if err := userDetailRepo.DB.Create(&UserDetail).Error; err != nil {
		return err
	}

	return nil
}

func (userDetailRepo *Repository) UpdateUserDetail(id int, UserDetail *eu.UserDetail) error {
	result := userDetailRepo.DB.Model(&UserDetail).Where("id = ?", id).Omit("UpdatedAt").Updates(&UserDetail)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user detail with id %d not found", id)
	}

	return nil
}

func (userDetailRepo *Repository) DeleteUserDetail(id int) error {
	if err := userDetailRepo.DB.Delete(&eu.UserDetail{}, id).Error; err != nil {
		return err
	}

	return nil
}

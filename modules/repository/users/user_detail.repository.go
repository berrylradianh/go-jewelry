package users

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (userDetailRepo *Repository) GetAllUserDetails() (*[]e.UserDetail, error) {
	var userDetails []e.UserDetail
	if err := userDetailRepo.DB.Preload("User", "deleted_at IS NULL").Find(&userDetails).Error; err != nil {
		return nil, err
	}

	return &userDetails, nil
}

func (userDetailRepo *Repository) GetUserDetailById(id int) (*e.UserDetail, error) {
	var UserDetail e.UserDetail
	if err := userDetailRepo.DB.Preload("User", "deleted_at IS NULL").First(&UserDetail, id).Error; err != nil {
		return nil, err
	}

	return &UserDetail, nil
}

func (userDetailRepo *Repository) CreateUserDetail(UserDetail *e.UserDetail) error {
	if err := userDetailRepo.DB.Create(&UserDetail).Error; err != nil {
		return err
	}

	return nil
}

func (userDetailRepo *Repository) UpdateUserDetail(id int, UserDetail *e.UserDetail) error {
	result := userDetailRepo.DB.Model(&UserDetail).Where("id = ?", id).Omit("UpdatedAt").Updates(&UserDetail)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (userDetailRepo *Repository) DeleteUserDetail(id int) error {
	if err := userDetailRepo.DB.Delete(&e.UserDetail{}, id).Error; err != nil {
		return err
	}

	return nil
}

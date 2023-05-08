package auth

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (authRepo *Repository) RegisterUser(user *e.User) error {
	if err := authRepo.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

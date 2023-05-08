package auth

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (authRepo *Repository) LoginUser(email string) (*e.User, error) {
	var user e.User
	if err := authRepo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

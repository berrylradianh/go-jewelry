package auth

import (
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
)

func (authRepo *Repository) RegisterUser(user *eu.User) error {
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	result := authRepo.DB.Create(&user)

	return result.Error
}

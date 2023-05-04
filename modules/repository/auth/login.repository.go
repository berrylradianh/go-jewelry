package auth

import (
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
)

func (authRepo *Repository) LoginUser(email string) (*eu.User, error) {
	var user eu.User
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	result := authRepo.DB.Where("email = ?", email).First(&user)

	return &user, result.Error
}

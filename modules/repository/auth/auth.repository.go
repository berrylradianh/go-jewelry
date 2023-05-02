package auth

import (
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (authRepo *Repository) LoginUser(email string) (*eu.User, error) {
	var user eu.User
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	result := authRepo.DB.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (authRepo *Repository) RegisterUser(user *eu.User) error {
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	result := authRepo.DB.Create(&user)

	return result.Error
}

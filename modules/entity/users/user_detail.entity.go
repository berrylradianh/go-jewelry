package users

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	*gorm.Model

	Name    string `json:"email" form:"email" validate:"required,email"`
	Address string `json:"address" form:"address" validate:"required"`
	Gender  string `json:"gender" form:"gender" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required"`
}

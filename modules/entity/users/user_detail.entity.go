package users

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	*gorm.Model

	Name    string `json:"name" form:"name" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Gender  string `json:"gender" form:"gender" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required"`
	User_id int    `json:"user_id,omitempty" form:"user_id" validate:"required"`
}

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

type UserDetailResponse struct {
	Name    string `json:"name,omitempty" form:"name"`
	Address string `json:"address,omitempty" form:"address"`
	Gender  string `json:"gender,omitempty" form:"gender"`
	Phone   string `json:"phone,omitempty" form:"phone"`
	User_id int    `json:"-"`
}

func (UserDetailResponse) TableName() string {
	return "user_details"
}

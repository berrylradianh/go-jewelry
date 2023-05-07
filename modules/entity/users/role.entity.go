package users

import (
	// eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model

	Name  string         `json:"name" form:"name" validate:"required"`
	Users []UserResponse `gorm:"foreignKey:Role_id"`
}

type RoleResponse struct {
	*gorm.Model `json:"-"`
	Name        string `json:"name" form:"name"`
	Users       []User `gorm:"foreignKey:Role_id" json:"-"`
}

func (RoleResponse) TableName() string {
	return "roles"
}

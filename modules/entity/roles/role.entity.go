package roles

import (
	ru "github.com/berrylradianh/go-jewelry/modules/response/users"
	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model

	Name string            `json:"name" form:"name" validate:"required"`
	User []ru.UserResponse `gorm:"foreignKey:Role_id" json:"users" form:"users"`
}

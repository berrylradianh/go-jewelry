package users

import (
	rr "github.com/berrylradianh/go-jewelry/modules/response/roles"
	ru "github.com/berrylradianh/go-jewelry/modules/response/users"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email       string                `json:"email" form:"email" validate:"required,email"`
	Password    string                `json:"password" form:"password" validate:"required"`
	User_detail ru.UserDetailResponse `gorm:"foreignKey:User_id" json:"user_detail" form:"user_detail"`
	Role_id     int                   `json:"role_id,omitempty" form:"role_id" validate:"required"`
	Role        rr.RoleResponse       `gorm:"foreignKey:ID" json:"role" form:"role"`
}

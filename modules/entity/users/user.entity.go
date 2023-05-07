package users

import (
	// er "github.com/berrylradianh/go-jewelry/modules/entity/roles"
	// et "github.com/berrylradianh/go-jewelry/modules/entity/transactions"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email       string             `json:"email" form:"email" validate:"required,email"`
	Password    string             `json:"password" form:"password" validate:"required"`
	User_detail UserDetailResponse `gorm:"foreignKey:User_id "`
	// Transaction []et.TransactionResponse `gorm:"foreignKey:User_id" json:"transaction" form:"transaction"`
	Role_id uint         `json:"role_id,omitempty" form:"role_id" validate:"required"`
	Role    RoleResponse `gorm:"foreignKey:Role_id"`
}

type UserResponse struct {
	*gorm.Model `json:"-"`
	Email       string       `json:"email,omitempty" form:"email"`
	Role_id     uint         `json:"-"`
	User_detail []UserDetail `gorm:"foreignKey:User_id" json:"-"`
}
type AuthResponse struct {
	ID    int    `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

func (UserResponse) TableName() string {
	return "users"
}

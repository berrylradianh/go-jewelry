package users

import (
	et "github.com/berrylradianh/go-jewelry/modules/entity/transactions"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email       string                   `json:"email" form:"email" validate:"required,email"`
	Password    string                   `json:"password" form:"password" validate:"required"`
	User_detail UserDetailResponse       `json:"user_detail" form:"user_detail"`
	Transaction []et.TransactionResponse `gorm:"foreignKey:User_id" json:"transaction" form:"transaction"`
	Role_id     int                      `json:"role_id,omitempty" form:"role_id" validate:"required"`
}

type UserResponse struct {
	ID      int    `json:"-"`
	Email   string `json:"email,omitempty" form:"email"`
	Role_id int    `json:"-"`
}
type AuthResponse struct {
	ID    int    `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

func (UserResponse) TableName() string {
	return "users"
}

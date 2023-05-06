package users

import (
	rt "github.com/berrylradianh/go-jewelry/modules/response/transactions"
	ru "github.com/berrylradianh/go-jewelry/modules/response/users"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email       string                   `json:"email" form:"email" validate:"required,email"`
	Password    string                   `json:"password" form:"password" validate:"required"`
	User_detail ru.UserDetailResponse    `gorm:"foreignKey:User_id" json:"user_detail" form:"user_detail"`
	Transaction []rt.TransactionResponse `gorm:"foreignKey:User_id" json:"transaction" form:"transaction"`
	Role_id     int                      `json:"role_id,omitempty" form:"role_id" validate:"required"`
}

package roles

// import (
// 	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
// 	"gorm.io/gorm"
// )

// type Role struct {
// 	*gorm.Model

// 	Name string            `json:"name" form:"name" validate:"required"`
// 	User []eu.UserResponse `gorm:"foreignKey:Role_id" json:"users" form:"users"`
// }

// type RoleResponse struct {
// 	ID   string `json:"-"`
// 	Name string `json:"name" form:"name"`
// }

// func (RoleResponse) TableName() string {
// 	return "roles"
// }

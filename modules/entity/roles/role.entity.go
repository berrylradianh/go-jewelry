package roles

import (
	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model

	Name string `json:"name" form:"name" validate:"required"`
}

package payments

import (
	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model

	Name string `json:"name" form:"name" validate:"required"`
}

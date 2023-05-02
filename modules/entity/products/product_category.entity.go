package products

import (
	"gorm.io/gorm"
)

type ProductCategory struct {
	*gorm.Model

	Name string `json:"name" form:"name" validate:"required"`
}

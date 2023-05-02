package products

import (
	"gorm.io/gorm"
)

type ProductMaterial struct {
	*gorm.Model

	Name string `json:"name" form:"name" validate:"required"`
}

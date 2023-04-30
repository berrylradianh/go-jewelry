package products

import (
	"gorm.io/gorm"
)

type ProductDescription struct {
	*gorm.Model

	Description string `json:"description" form:"description" validate:"required"`
}

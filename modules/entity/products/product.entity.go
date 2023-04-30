package products

import (
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model

	Name      string `json:"name" form:"name" validate:"required"`
	Price     string `json:"price" form:"price" validate:"required"`
	Stock     int    `json:"stock" form:"stock" validate:"required"`
	Image_url string `json:"image_url" form:"image_url" validate:"required"`
}

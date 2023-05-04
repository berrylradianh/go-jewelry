package products

import (
	pr "github.com/berrylradianh/go-jewelry/modules/response/products"
	"gorm.io/gorm"
)

type ProductDescription struct {
	*gorm.Model

	Description string             `json:"description" form:"description" validate:"required"`
	Products    pr.ProductResponse `gorm:"foreignKey:Product_description_id" json:"products" form:"products"`
}

package product_categories

import (
	pr "github.com/berrylradianh/go-jewelry/modules/response/products"
	"gorm.io/gorm"
)

type ProductCategory struct {
	*gorm.Model

	Name    string `json:"name" form:"name" validate:"required"`
	Product []pr.ProductResponse
}

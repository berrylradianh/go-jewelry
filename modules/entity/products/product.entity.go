package products

import (
	pc "github.com/berrylradianh/go-jewelry/modules/response/products"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model

	Name                   string                        `json:"name" form:"name" validate:"required"`
	Price                  float64                       `json:"price" form:"price" validate:"required"`
	Stock                  int                           `json:"stock" form:"stock" validate:"required"`
	Image_url              string                        `json:"image_url" form:"image_url" validate:"required"`
	Product_category_id    int                           `json:"product_category_id" form:"product_category_id" validate:"required"`
	Product_category       pc.ProductCategoryResponse    `gorm:"foreignKey:ID" json:"product_category" form:"product_category"`
	Product_material_id    int                           `json:"product_material_id" form:"product_material_id" validate:"required"`
	Product_material       pc.ProductMaterialResponse    `gorm:"foreignKey:ID" json:"product_material" form:"product_material"`
	Product_description_id int                           `json:"product_description_id" form:"product_description_id" validate:"required"`
	Product_description    pc.ProductDescriptionResponse `gorm:"foreignKey:ID" json:"product_description" form:"product_description"`
}

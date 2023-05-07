package products

import (
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model

	Name      string  `json:"name,omitempty" form:"name"`
	Price     float64 `json:"price,omitempty" form:"price"`
	Stock     int     `json:"stock,omitempty" form:"stock"`
	Image_url string  `json:"image_url,omitempty" form:"image_url"`

	Product_category_id uint                    `json:"product_category_id,omitempty" form:"product_category_id"`
	Product_category    ProductCategoryResponse `gorm:"foreignKey:Product_category_id"`
	Product_material_id uint                    `json:"product_material_id,omitempty" form:"product_material_id"`
	Product_material    ProductMaterialResponse `gorm:"foreignKey:Product_material_id"`

	// Product_material_id uint                    `json:"product_material_id,omitempty" form:"product_material_id"`
	// Product_material    ProductMaterialResponse `gorm:"foreignKey:Product_material_id"`
	// Product_description ProductDescriptionResponse     `gorm:"foreignKey:Product_description_id" json:"product_description" form:"product_description"`
	// Transaction_detail  []et.TransactionDetailResponse `gorm:"foreignKey:Transaction_id" json:"transaction_detail,omitempty" form:"transaction_detail"`
}

type ProductResponse struct {
	ID                  uint    `json:"id,omitempty" form:"id"`
	Name                string  `json:"name,omitempty" form:"name"`
	Price               float64 `json:"price,omitempty" form:"price"`
	Stock               int     `json:"stock,omitempty" form:"stock"`
	Image_url           string  `json:"image_url,omitempty" form:"image_url"`
	Product_category_id uint    `json:"-"`
	Product_material_id uint    `json:"-"`
	// Product_description_id int     `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

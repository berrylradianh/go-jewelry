package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model

	Name      string  `json:"name,omitempty" form:"name"`
	Price     float64 `json:"price,omitempty" form:"price"`
	Stock     int     `json:"stock,omitempty" form:"stock"`
	Image_url string  `json:"image_url,omitempty" form:"image_url"`

	Product_category_id uint                        `json:"product_category_id,omitempty" form:"product_category_id"`
	Product_category    ProductCategoryResponse     `gorm:"foreignKey:Product_category_id"`
	Cart_id             uint                        `json:"cart_id,omitempty" form:"cart_id"`
	Cart                CartResponse                `gorm:"foreignKey:Cart_id"`
	Product_material_id uint                        `json:"product_material_id,omitempty" form:"product_material_id"`
	Product_material    ProductMaterialResponse     `gorm:"foreignKey:Product_material_id"`
	Product_description ProductDescriptionResponse  `gorm:"foreignKey:Product_id "`
	Transaction_details []TransactionDetailResponse `gorm:"foreignKey:Product_id"`
}

type ProductResponse struct {
	*gorm.Model         `json:"-"`
	Name                string               `json:"name,omitempty" form:"name"`
	Price               float64              `json:"price,omitempty" form:"price"`
	Stock               int                  `json:"stock,omitempty" form:"stock"`
	Image_url           string               `json:"image_url,omitempty" form:"image_url"`
	Product_category_id uint                 `json:"-"`
	Cart_id             uint                 `json:"-"`
	Product_material_id uint                 `json:"-"`
	Product_description []ProductDescription `gorm:"foreignKey:Product_id  " json:"-"`
	Transaction_details []TransactionDetail  `gorm:"foreignKey:Product_id" json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

package products

import (
	"gorm.io/gorm"
)

type ProductMaterial struct {
	*gorm.Model

	Name     string            `json:"name" form:"name" validate:"required"`
	Products []ProductResponse `gorm:"foreignKey:Product_material_id"`
}

type ProductMaterialResponse struct {
	*gorm.Model `json:"-"`
	Name        string    `json:"name" form:"name"`
	Products    []Product `gorm:"foreignKey:Product_material_id" json:"-"`
}

func (ProductMaterialResponse) TableName() string {
	return "product_materials"
}

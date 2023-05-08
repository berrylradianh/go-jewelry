package entity

import (
	"gorm.io/gorm"
)

type ProductDescription struct {
	*gorm.Model

	Description string          `json:"description" form:"description" validate:"required"`
	Product_id  uint            `json:"product_id,omitempty" form:"product_id" validate:"required"`
	Product     ProductResponse `gorm:"foreignKey:Product_id"`
}

type ProductDescriptionResponse struct {
	*gorm.Model `json:"-"`
	Description string `json:"description,omitempty" form:"description"`
	Product_id  uint   `json:"-"`
}

func (ProductDescriptionResponse) TableName() string {
	return "product_descriptions"
}

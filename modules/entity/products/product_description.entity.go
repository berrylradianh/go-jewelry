package products

import (
	"gorm.io/gorm"
)

type ProductDescription struct {
	*gorm.Model

	Description string `json:"description" form:"description" validate:"required"`
	Product_id  int    `json:"product_id,omitempty" form:"product_id" validate:"required"`
	// Product     ProductResponse `gorm:"foreignKey:Product_description_id" json:"products" form:"products"`
}

type ProductDescriptionResponse struct {
	ID                     int    `json:"-"`
	Description            string `json:"description" form:"description"`
	Product_description_id int    `json:"-"`
}

func (ProductDescriptionResponse) TableName() string {
	return "product_descriptions"
}

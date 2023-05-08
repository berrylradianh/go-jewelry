package entity

import (
	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	Product_id uint            `json:"product_id,omitempty" form:"product_id"`
	Product    ProductResponse `gorm:"foreignKey:Product_id"`
	Qty        float64         `json:"qty" form:"qty" validate:"required"`
	Price      float64         `json:"price" form:"price"`
}

type CartResponse struct {
	*gorm.Model `json:"-"`
	Qty         float64 `json:"qty" form:"qty"`
	Price       float64 `json:"price" form:"price"`
	Product_id  uint    `json:"-"`
}

func (CartResponse) TableName() string {
	return "carts"
}

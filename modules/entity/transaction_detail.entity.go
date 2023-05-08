package entity

import (
	"gorm.io/gorm"
)

type TransactionDetail struct {
	*gorm.Model

	Qty            int16               `json:"qty" form:"qty" validate:"required"`
	Price          float64             `json:"price" form:"price"`
	Transaction_id uint                `json:"transaction_id,omitempty" form:"transaction_id" validate:"required"`
	Transaction    TransactionResponse `gorm:"foreignKey:Transaction_id"`
	Product_id     uint                `json:"product_id,omitempty" form:"product_id"`
	Product        ProductResponse     `gorm:"foreignKey:Product_id"`
}

type TransactionDetailResponse struct {
	*gorm.Model    `json:"-"`
	Qty            int16   `json:"qty" form:"qty"`
	Price          float64 `json:"price" form:"price"`
	Transaction_id int     `json:"-"`
	Product_id     uint    `json:"-"`
}

func (TransactionDetailResponse) TableName() string {
	return "transaction_details"
}

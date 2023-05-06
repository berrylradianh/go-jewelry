package transactions

import (
	"gorm.io/gorm"
)

type TransactionDetail struct {
	*gorm.Model

	Qty            int16   `json:"qty" form:"qty" validate:"required"`
	Price          float64 `json:"price" form:"price" validate:"required"`
	Transaction_id int     `json:"transaction_id,omitempty" form:"transaction_id" validate:"required"`
	Product_id     int     `json:"product_id,omitempty" form:"product_id" validate:"required"`
}

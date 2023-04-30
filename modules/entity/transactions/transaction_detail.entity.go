package transactions

import (
	"gorm.io/gorm"
)

type TransactionDetail struct {	
	*gorm.Model

	Qty        string `json:"qty" form:"qty" validate:"required"`
	TotalPrice string `json:"total_price" form:"total_price" validate:"required"`
	Image_url  string `json:"image_url" form:"image_url" validate:"required"`
}

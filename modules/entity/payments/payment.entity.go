package payments

import (
	rt "github.com/berrylradianh/go-jewelry/modules/response/transactions"
	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model

	Name        string                   `json:"name" form:"name" validate:"required"`
	Transaction []rt.TransactionResponse `gorm:"foreignKey:Payment_id" json:"transaction" form:"transaction"`
}

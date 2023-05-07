package payments

import (
	// et "github.com/berrylradianh/go-jewelry/modules/entity/transactions"
	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model

	Name string `json:"name" form:"name" validate:"required"`
	// Transaction []et.TransactionResponse `gorm:"foreignKey:Payment_id" json:"transaction" form:"transaction"`
}

type PaymentResponse struct {
	*gorm.Model `json:"-"`
	Name        string `json:"name,omitempty" form:"name"`
}

func (PaymentResponse) TableName() string {
	return "payments"
}

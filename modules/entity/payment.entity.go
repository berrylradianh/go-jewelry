package entity

import (
	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model

	Name string `json:"name" form:"name" validate:"required"`
	// Transaction []et.TransactionResponse `gorm:"foreignKey:Payment_id" json:"transaction" form:"transaction"`
	Transaction []TransactionResponse `gorm:"foreignKey:Payment_id"`
}

type PaymentResponse struct {
	*gorm.Model `json:"-"`
	Name        string        `json:"name,omitempty" form:"name"`
	Transaction []Transaction `gorm:"foreignKey:Payment_id" json:"-"`
}

func (PaymentResponse) TableName() string {
	return "payments"
}

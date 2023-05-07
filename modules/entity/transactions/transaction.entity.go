package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model

	Date            time.Time `json:"date" form:"date"`
	Status          string    `json:"status" form:"status"` // Verified - Process - Canceled
	Image_proof_url string    `json:"image_proof_url" form:"image_proof_url" validate:"required"`
	// User_id            int                         `json:"user_id,omitempty" form:"user_id" validate:"required"`
	// Payment_id         int                         `json:"payment_id,omitempty" form:"payment" validate:"required"`
	Transaction_detail []TransactionDetailResponse `gorm:"foreignKey:Transaction_id"`
}

type TransactionResponse struct {
	*gorm.Model     `json:"-"`
	Date            time.Time `json:"date" form:"date"`
	Status          string    `json:"status" form:"status"`
	Image_proof_url string    `json:"image_proof_url" form:"image_proof_url"`
	// User_id            int                 `json:"-" form:"-"`
	// Payment_id         int                 `json:"-" form:"-"`
	Transaction_detail []TransactionDetail `gorm:"foreignKey:Transaction_id" json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}

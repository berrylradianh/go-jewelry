package transactions

import (
	"time"

	rt "github.com/berrylradianh/go-jewelry/modules/response/transactions"

	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model

	Date               time.Time                      `json:"date" form:"date"`
	Status             string                         `json:"status" form:"status"` // Verified - Process - Canceled
	Image_proof_url    string                         `json:"image_proof_url" form:"image_proof_url" validate:"required"`
	User_id            int                            `json:"user_id,omitempty" form:"user_id" validate:"required"`
	Payment_id         int                            `json:"payment_id,omitempty" form:"payment" validate:"required"`
	Transaction_detail []rt.TransactionDetailResponse `gorm:"foreignKey:Transaction_id" json:"transaction_detail" form:"transaction_detail"`
}

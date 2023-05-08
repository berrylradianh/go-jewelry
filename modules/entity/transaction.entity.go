package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model

	Date                time.Time                   `json:"date" form:"date"`
	Status              string                      `json:"status" form:"status"` // Verified - Process - Canceled
	Image_proof_url     string                      `json:"image_proof_url" form:"image_proof_url" validate:"required"`
	Product_category_id uint                        `json:"product_category_id,omitempty" form:"product_category_id"`
	User_id             uint                        `json:"user_id,omitempty" form:"user_id" validate:"required"`
	User                UserResponse                `gorm:"foreignKey:User_id"`
	Payment_id          uint                        `json:"payment_id,omitempty" form:"user_id" validate:"required"`
	Payment             PaymentResponse             `gorm:"foreignKey:Payment_id"`
	Transaction_details []TransactionDetailResponse `gorm:"foreignKey:Transaction_id"`
}

type TransactionResponse struct {
	*gorm.Model         `json:"-"`
	Date                time.Time           `json:"date" form:"date"`
	Status              string              `json:"status" form:"status"`
	Image_proof_url     string              `json:"image_proof_url" form:"image_proof_url"`
	User_id             uint                `json:"-"`
	Payment_id          uint                `json:"-" form:"-"`
	Transaction_details []TransactionDetail `gorm:"foreignKey:Transaction_id" json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}

package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	*gorm.Model

	Date   time.Time `json:"date" form:"date" validate:"required"`
	Status string    `json:"status" form:"status" validate:"required"`
}

package entity

import (
	"gorm.io/gorm"
)

type Cart struct {
	*gorm.Model
	Products []ProductResponse `gorm:"foreignKey:Cart_id"`
}

type CartResponse struct {
	*gorm.Model `json:"-"`
	Products    []Product `gorm:"foreignKey:Cart_id" json:"-"`
}

func (CartResponse) TableName() string {
	return "carts"
}

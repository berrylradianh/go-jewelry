package carts

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (cartRepo *Repository) GetAllCarts() (*[]e.Cart, error) {
	var carts []e.Cart
	if err := cartRepo.DB.Preload("Product", "deleted_at IS NULL").Find(&carts).Error; err != nil {
		return nil, err
	}

	return &carts, nil
}

func (cartRepo *Repository) GetCartById(id int) (*e.Cart, error) {
	var cart e.Cart
	if err := cartRepo.DB.Preload("Product", "deleted_at IS NULL").First(&cart, id).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (cartRepo *Repository) FindProduct(id int) (*e.Product, error) {
	var product e.Product
	if err := cartRepo.DB.First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (cartRepo *Repository) CreateCart(cart *e.Cart) error {
	if err := cartRepo.DB.Create(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (cartRepo *Repository) UpdateCart(id int, cart *e.Cart) error {
	result := cartRepo.DB.Model(&cart).Where("id = ?", id).Omit("UpdatedAt").Updates(&cart)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (cartRepo *Repository) DeleteCart(id int) error {
	if err := cartRepo.DB.Delete(&e.Cart{}, id).Error; err != nil {
		return err
	}

	return nil
}

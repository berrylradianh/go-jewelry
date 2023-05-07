package products

import (
	"fmt"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
)

func (productRepo *Repository) GetAllProducts() (*[]ep.Product, error) {
	var products []ep.Product
	if err := productRepo.DB.Preload("Product_category", "deleted_at IS NULL").Preload("Product_material", "deleted_at IS NULL").Preload("Product_description", "deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (productRepo *Repository) GetProductById(id int) (*ep.Product, error) {
	var product ep.Product
	if err := productRepo.DB.Preload("Product_category", "deleted_at IS NULL").Preload("Product_material", "deleted_at IS NULL").First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (productRepo *Repository) CreateProduct(product *ep.Product) error {
	if err := productRepo.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (productRepo *Repository) UpdateProduct(id int, product *ep.Product) error {
	result := productRepo.DB.Model(&product).Where("id = ?", id).Omit("UpdatedAt").Updates(&product)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product category with id %d not found", id)
	}

	return nil
}

func (productRepo *Repository) DeleteProduct(id int) error {
	if err := productRepo.DB.Delete(&ep.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

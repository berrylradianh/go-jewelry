package products

import (
	"fmt"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
)

func (productDescriptionRepo *Repository) GetAllProductDescriptions() (*[]ep.ProductDescription, error) {
	var productDescriptions []ep.ProductDescription
	if err := productDescriptionRepo.DB.Preload("Product", "deleted_at IS NULL").Find(&productDescriptions).Error; err != nil {
		return nil, err
	}

	return &productDescriptions, nil
}

func (productDescriptionRepo *Repository) GetProductDescriptionById(id int) (*ep.ProductDescription, error) {
	var productDescription ep.ProductDescription
	if err := productDescriptionRepo.DB.Preload("Product", "deleted_at IS NULL").First(&productDescription, id).Error; err != nil {
		return nil, err
	}

	return &productDescription, nil
}

func (productDescriptionRepo *Repository) CreateProductDescription(productDescription *ep.ProductDescription) error {
	if err := productDescriptionRepo.DB.Create(&productDescription).Error; err != nil {
		return err
	}

	return nil
}

func (productDescriptionRepo *Repository) UpdateProductDescription(id int, productDescription *ep.ProductDescription) error {
	result := productDescriptionRepo.DB.Model(&productDescription).Where("id = ?", id).Omit("UpdatedAt").Updates(&productDescription)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product Description with id %d not found", id)
	}

	return nil
}

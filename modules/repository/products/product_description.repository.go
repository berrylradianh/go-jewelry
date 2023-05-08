package products

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productDescriptionRepo *Repository) GetAllProductDescriptions() (*[]e.ProductDescription, error) {
	var productDescriptions []e.ProductDescription
	if err := productDescriptionRepo.DB.Preload("Product", "deleted_at IS NULL").Find(&productDescriptions).Error; err != nil {
		return nil, err
	}

	return &productDescriptions, nil
}

func (productDescriptionRepo *Repository) GetProductDescriptionById(id int) (*e.ProductDescription, error) {
	var productDescription e.ProductDescription
	if err := productDescriptionRepo.DB.Preload("Product", "deleted_at IS NULL").First(&productDescription, id).Error; err != nil {
		return nil, err
	}

	return &productDescription, nil
}

func (productDescriptionRepo *Repository) CreateProductDescription(productDescription *e.ProductDescription) error {
	if err := productDescriptionRepo.DB.Create(&productDescription).Error; err != nil {
		return err
	}

	return nil
}

func (productDescriptionRepo *Repository) UpdateProductDescription(id int, productDescription *e.ProductDescription) error {
	result := productDescriptionRepo.DB.Model(&productDescription).Where("id = ?", id).Omit("UpdatedAt").Updates(&productDescription)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

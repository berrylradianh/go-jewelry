package products

import (
	"fmt"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
)

func (productCategoryRepo *Repository) GetAllProductCategories() (*[]ep.ProductCategory, error) {
	var productCategories []ep.ProductCategory
	if err := productCategoryRepo.DB.Preload("Products", "deleted_at IS NULL").Find(&productCategories).Error; err != nil {
		return nil, err
	}

	return &productCategories, nil
}

func (productCategoryRepo *Repository) GetProductCategoryById(id int) (*ep.ProductCategory, error) {
	var productCategory ep.ProductCategory
	if err := productCategoryRepo.DB.Preload("Products", "deleted_at IS NULL").First(&productCategory, id).Error; err != nil {
		return nil, err
	}

	return &productCategory, nil
}

func (productCategoryRepo *Repository) CreateProductCategory(productCategory *ep.ProductCategory) error {
	if err := productCategoryRepo.DB.Create(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (productCategoryRepo *Repository) UpdateProductCategory(id int, productCategory *ep.ProductCategory) error {
	result := productCategoryRepo.DB.Model(&productCategory).Where("id = ?", id).Omit("UpdatedAt").Updates(&productCategory)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product category with id %d not found", id)
	}

	return nil
}

func (productCategoryRepo *Repository) DeleteProductCategory(id int) error {
	if err := productCategoryRepo.DB.Delete(&ep.ProductCategory{}, id).Error; err != nil {
		return err
	}

	return nil
}

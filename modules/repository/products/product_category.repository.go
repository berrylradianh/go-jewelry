package products

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productCategoryRepo *Repository) GetAllProductCategories() (*[]e.ProductCategory, error) {
	var productCategories []e.ProductCategory
	if err := productCategoryRepo.DB.Preload("Products", "deleted_at IS NULL").Find(&productCategories).Error; err != nil {
		return nil, err
	}

	return &productCategories, nil
}

func (productCategoryRepo *Repository) GetProductCategoryById(id int) (*e.ProductCategory, error) {
	var productCategory e.ProductCategory
	if err := productCategoryRepo.DB.Preload("Products", "deleted_at IS NULL").First(&productCategory, id).Error; err != nil {
		return nil, err
	}

	return &productCategory, nil
}

func (productCategoryRepo *Repository) CreateProductCategory(productCategory *e.ProductCategory) error {
	if err := productCategoryRepo.DB.Create(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (productCategoryRepo *Repository) UpdateProductCategory(id int, productCategory *e.ProductCategory) error {
	result := productCategoryRepo.DB.Model(&productCategory).Where("id = ?", id).Omit("UpdatedAt").Updates(&productCategory)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (productCategoryRepo *Repository) DeleteProductCategory(id int) error {
	if err := productCategoryRepo.DB.Delete(&e.ProductCategory{}, id).Error; err != nil {
		return err
	}

	return nil
}

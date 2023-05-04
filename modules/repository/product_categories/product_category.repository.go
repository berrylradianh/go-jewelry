package product_categories

import (
	"fmt"

	epc "github.com/berrylradianh/go-jewelry/modules/entity/product_categories"
	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (productCategoryRepo *Repository) GetAllProductCategories() (*[]epc.ProductCategory, error) {
	var productCategories []epc.ProductCategory
	if err := productCategoryRepo.DB.Preload("Products", "deleted_at IS NULL").Find(&productCategories).Error; err != nil {
		return nil, err
	}

	return &productCategories, nil
}

func (productCategoryRepo *Repository) GetProductCategoryById(id int) (*epc.ProductCategory, error) {
	var productCategory epc.ProductCategory
	if err := productCategoryRepo.DB.Preload("Products", "deleted_at IS NULL").First(&productCategory, id).Error; err != nil {
		return nil, err
	}

	return &productCategory, nil
}

func (productCategoryRepo *Repository) CreateProductCategory(productCategory *epc.ProductCategory) error {
	if err := productCategoryRepo.DB.Create(&productCategory).Error; err != nil {
		return err
	}

	return nil
}

func (productCategoryRepo *Repository) UpdateProductCategory(id int, productCategory *epc.ProductCategory) error {
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
	if err := productCategoryRepo.DB.Where("product_category_id = ?", id).Delete(&ep.Product{}).Error; err != nil {
		return err
	}

	if err := productCategoryRepo.DB.Delete(&epc.ProductCategory{}, id).Error; err != nil {
		return err
	}

	return nil
}

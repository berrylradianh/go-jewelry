package products

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productRepo *Repository) GetAllProducts() (*[]e.Product, error) {
	var products []e.Product
	if err := productRepo.DB.Preload("Product_category", "deleted_at IS NULL").Preload("Product_material", "deleted_at IS NULL").Preload("Product_description", "deleted_at IS NULL").Preload("Transaction_details", "deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (productRepo *Repository) GetProductById(id int) (*e.Product, error) {
	var product e.Product
	if err := productRepo.DB.Preload("Product_category", "deleted_at IS NULL").Preload("Product_material", "deleted_at IS NULL").Preload("Product_description", "deleted_at IS NULL").Preload("Transaction_details", "deleted_at IS NULL").First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (productRepo *Repository) CreateProduct(product *e.Product) error {
	if err := productRepo.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (productRepo *Repository) UpdateProduct(id int, product *e.Product) error {
	result := productRepo.DB.Model(&product).Where("id = ?", id).Omit("UpdatedAt").Updates(&product)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (productRepo *Repository) DeleteProduct(id int) error {
	if err := productRepo.DB.Where("product_id   = ?", id).Delete(&e.ProductDescription{}).Error; err != nil {
		return err
	}

	if err := productRepo.DB.Delete(&e.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (productRepo *Repository) SortProductByNameASC() (*[]e.Product, error) {
	var products []e.Product
	if err := productRepo.DB.Order("name").Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (productRepo *Repository) SortProductByNameDESC() (*[]e.Product, error) {
	var products []e.Product
	if err := productRepo.DB.Order("name DESC").Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (productRepo *Repository) SortProductByDateASC() (*[]e.Product, error) {
	var products []e.Product
	if err := productRepo.DB.Order("created_at").Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (productRepo *Repository) SortProductByDateDESC() (*[]e.Product, error) {
	var products []e.Product
	if err := productRepo.DB.Order("created_at DESC").Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

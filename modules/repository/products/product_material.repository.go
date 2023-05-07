package products

import (
	"fmt"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
)

func (productMaterialRepo *Repository) GetAllProductMaterials() (*[]ep.ProductMaterial, error) {
	var productMaterials []ep.ProductMaterial
	if err := productMaterialRepo.DB.Preload("Products", "deleted_at IS NULL").Find(&productMaterials).Error; err != nil {
		return nil, err
	}

	return &productMaterials, nil
}

func (productMaterialRepo *Repository) GetProductMaterialById(id int) (*ep.ProductMaterial, error) {
	var productMaterial ep.ProductMaterial
	if err := productMaterialRepo.DB.Preload("Products", "deleted_at IS NULL").First(&productMaterial, id).Error; err != nil {
		return nil, err
	}

	return &productMaterial, nil
}

func (productMaterialRepo *Repository) CreateProductMaterial(productMaterial *ep.ProductMaterial) error {
	if err := productMaterialRepo.DB.Create(&productMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (productMaterialRepo *Repository) UpdateProductMaterial(id int, productMaterial *ep.ProductMaterial) error {
	result := productMaterialRepo.DB.Model(&productMaterial).Where("id = ?", id).Omit("UpdatedAt").Updates(&productMaterial)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("product Material with id %d not found", id)
	}

	return nil
}

func (productMaterialRepo *Repository) DeleteProductMaterial(id int) error {
	if err := productMaterialRepo.DB.Delete(&ep.ProductMaterial{}, id).Error; err != nil {
		return err
	}

	return nil
}

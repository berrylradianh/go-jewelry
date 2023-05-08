package products

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productMaterialRepo *Repository) GetAllProductMaterials() (*[]e.ProductMaterial, error) {
	var productMaterials []e.ProductMaterial
	if err := productMaterialRepo.DB.Preload("Products", "deleted_at IS NULL").Find(&productMaterials).Error; err != nil {
		return nil, err
	}

	return &productMaterials, nil
}

func (productMaterialRepo *Repository) GetProductMaterialById(id int) (*e.ProductMaterial, error) {
	var productMaterial e.ProductMaterial
	if err := productMaterialRepo.DB.Preload("Products", "deleted_at IS NULL").First(&productMaterial, id).Error; err != nil {
		return nil, err
	}

	return &productMaterial, nil
}

func (productMaterialRepo *Repository) CreateProductMaterial(productMaterial *e.ProductMaterial) error {
	if err := productMaterialRepo.DB.Create(&productMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (productMaterialRepo *Repository) UpdateProductMaterial(id int, productMaterial *e.ProductMaterial) error {
	result := productMaterialRepo.DB.Model(&productMaterial).Where("id = ?", id).Omit("UpdatedAt").Updates(&productMaterial)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (productMaterialRepo *Repository) DeleteProductMaterial(id int) error {
	if err := productMaterialRepo.DB.Delete(&e.ProductMaterial{}, id).Error; err != nil {
		return err
	}

	return nil
}

package product_categories

import (
	epc "github.com/berrylradianh/go-jewelry/modules/entity/product_categories"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (productCategoryRepo *Repository) GetAllProductCategories() (*[]epc.ProductCategory, error) {
	var productCategories []epc.ProductCategory
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := productCategoryRepo.DB.Preload("product_categories", "deleted_at IS NULL").Find(&products)
	result := productCategoryRepo.DB.Find(&productCategories)

	return &productCategories, result.Error
}

func (productCategoryRepo *Repository) GetProductCategoryById(id int) (*epc.ProductCategory, error) {
	var productCategory epc.ProductCategory
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := productCategoryRepo.DB.Preload("product_categories", "deleted_at IS NULL").First(&product, id)
	result := productCategoryRepo.DB.First(&productCategory, id)

	return &productCategory, result.Error
}

func (productCategoryRepo *Repository) CreateProductCategory(productCategory *epc.ProductCategory) error {
	result := productCategoryRepo.DB.Create(&productCategory)
	return result.Error
}

func (productCategoryRepo *Repository) UpdateProductCategory(id int, productCategory *epc.ProductCategory) int64 {
	result := productCategoryRepo.DB.Model(&productCategory).Where("id = ?", id).Omit("UpdatedAt").Updates(&productCategory)
	return result.RowsAffected
}

func (productCategoryRepo *Repository) DeleteProductCategory(id int) error {
	result := productCategoryRepo.DB.Delete(&epc.ProductCategory{}, id)
	return result.Error
}

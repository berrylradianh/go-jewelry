package products

import (
	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (productRepo *Repository) GetAllProducts() (*[]ep.Product, error) {
	var products []ep.Product
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := productRepo.DB.Preload("product_categories", "deleted_at IS NULL").Find(&products)
	result := productRepo.DB.Find(&products)

	return &products, result.Error
}

func (productRepo *Repository) GetProductById(id int) (*ep.Product, error) {
	var product ep.Product
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := productRepo.DB.Preload("product_categories", "deleted_at IS NULL").First(&product, id)
	result := productRepo.DB.First(&product, id)

	return &product, result.Error
}

func (productRepo *Repository) CreateProduct(product *ep.Product) error {
	result := productRepo.DB.Create(&product)
	return result.Error
}

func (productRepo *Repository) UpdateProduct(id int, product *ep.Product) int64 {
	result := productRepo.DB.Model(&product).Where("id = ?", id).Omit("UpdatedAt").Updates(&product)
	return result.RowsAffected
}

func (productRepo *Repository) DeleteProduct(id int) error {
	result := productRepo.DB.Delete(&ep.Product{}, id)
	return result.Error
}

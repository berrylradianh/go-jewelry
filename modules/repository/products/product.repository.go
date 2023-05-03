package products

import (
	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (authRepo *Repository) GetAllProducts() (*[]ep.Product, error) {
	var products []ep.Product
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := authRepo.DB.Preload("product_categories", "deleted_at IS NULL").Find(&products)
	result := authRepo.DB.Find(&products)

	return &products, result.Error
}

func (authRepo *Repository) GetProductById(id int) (*ep.Product, error) {
	var product ep.Product
	// result := repo.DB.Preload("Blogs", "deleted_at IS NULL").Find(&users)
	// result := authRepo.DB.Preload("product_categories", "deleted_at IS NULL").First(&product, id)
	result := authRepo.DB.First(&product, id)

	return &product, result.Error
}

func (authRepo *Repository) CreateProduct(product *ep.Product) error {
	result := authRepo.DB.Create(&product)
	return result.Error
}

func (authRepo *Repository) UpdateProduct(id int, product *ep.Product) *gorm.DB {
	result := authRepo.DB.Model(&product).Where("id = ?", id).Omit("UpdatedAt").Updates(&product)
	return result
}

func (authRepo *Repository) DeleteProduct(id int) error {
	result := authRepo.DB.Delete(&ep.Product{}, id)
	return result.Error
}

package products

import (
	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
	rp "github.com/berrylradianh/go-jewelry/modules/repository/products"
	"gorm.io/gorm"
)

type Usecase struct {
	Repository rp.Repository
}

func (productUsecase *Usecase) GetAllProducts() (*[]ep.Product, error) {
	products, err := productUsecase.Repository.GetAllProducts()
	return products, err
}

func (productUsecase *Usecase) GetProductById(id int) (*ep.Product, error) {
	product, err := productUsecase.Repository.GetProductById(id)
	return product, err
}

func (productUsecase *Usecase) CreateProduct(product *ep.Product) error {
	err := productUsecase.Repository.CreateProduct(product)
	return err
}

func (productUsecase *Usecase) UpdateProduct(id int, product *ep.Product) *gorm.DB {
	result := productUsecase.Repository.UpdateProduct(id, product)
	return result
}

func (productUsecase *Usecase) DeleteProduct(id int) error {
	err := productUsecase.Repository.DeleteProduct(id)
	return err
}

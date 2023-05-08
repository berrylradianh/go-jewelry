package products

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productUsecase *Usecase) GetAllProducts() (*[]e.Product, error) {
	products, err := productUsecase.Repository.GetAllProducts()
	return products, err
}

func (productUsecase *Usecase) GetProductById(id int) (*e.Product, error) {
	product, err := productUsecase.Repository.GetProductById(id)
	return product, err
}

func (productUsecase *Usecase) CreateProduct(product *e.Product) error {
	err := productUsecase.Repository.CreateProduct(product)
	return err
}

func (productUsecase *Usecase) UpdateProduct(id int, product *e.Product) error {
	result := productUsecase.Repository.UpdateProduct(id, product)
	return result
}

func (productUsecase *Usecase) DeleteProduct(id int) error {
	err := productUsecase.Repository.DeleteProduct(id)
	return err
}

package products

import (
	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
)

func (productDescriptionUsecase *Usecase) GetAllProductDescriptions() (*[]ep.ProductDescription, error) {
	productDescriptions, err := productDescriptionUsecase.Repository.GetAllProductDescriptions()
	return productDescriptions, err
}

func (productDescriptionUsecase *Usecase) GetProductDescriptionById(id int) (*ep.ProductDescription, error) {
	productDescription, err := productDescriptionUsecase.Repository.GetProductDescriptionById(id)
	return productDescription, err
}

func (productDescriptionUsecase *Usecase) CreateProductDescription(productDescription *ep.ProductDescription) error {
	err := productDescriptionUsecase.Repository.CreateProductDescription(productDescription)
	return err
}

func (productDescriptionUsecase *Usecase) UpdateProductDescription(id int, productDescription *ep.ProductDescription) error {
	err := productDescriptionUsecase.Repository.UpdateProductDescription(id, productDescription)
	return err
}

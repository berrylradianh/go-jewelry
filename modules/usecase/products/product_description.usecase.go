package products

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productDescriptionUsecase *Usecase) GetAllProductDescriptions() (*[]e.ProductDescription, error) {
	productDescriptions, err := productDescriptionUsecase.Repository.GetAllProductDescriptions()
	return productDescriptions, err
}

func (productDescriptionUsecase *Usecase) GetProductDescriptionById(id int) (*e.ProductDescription, error) {
	productDescription, err := productDescriptionUsecase.Repository.GetProductDescriptionById(id)
	return productDescription, err
}

func (productDescriptionUsecase *Usecase) CreateProductDescription(productDescription *e.ProductDescription) error {
	err := productDescriptionUsecase.Repository.CreateProductDescription(productDescription)
	return err
}

func (productDescriptionUsecase *Usecase) UpdateProductDescription(id int, productDescription *e.ProductDescription) error {
	err := productDescriptionUsecase.Repository.UpdateProductDescription(id, productDescription)
	return err
}

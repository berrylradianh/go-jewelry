package products

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productCategoryUsecase *Usecase) GetAllProductCategories() (*[]e.ProductCategory, error) {
	productCategories, err := productCategoryUsecase.Repository.GetAllProductCategories()
	return productCategories, err
}

func (productCategoryUsecase *Usecase) GetProductCategoryById(id int) (*e.ProductCategory, error) {
	productCategory, err := productCategoryUsecase.Repository.GetProductCategoryById(id)
	return productCategory, err
}

func (productCategoryUsecase *Usecase) CreateProductCategory(productCategory *e.ProductCategory) error {
	err := productCategoryUsecase.Repository.CreateProductCategory(productCategory)
	return err
}

func (productCategoryUsecase *Usecase) UpdateProductCategory(id int, productCategory *e.ProductCategory) error {
	err := productCategoryUsecase.Repository.UpdateProductCategory(id, productCategory)
	return err
}

func (productCategoryUsecase *Usecase) DeleteProductCategory(id int) error {
	err := productCategoryUsecase.Repository.DeleteProductCategory(id)
	return err
}

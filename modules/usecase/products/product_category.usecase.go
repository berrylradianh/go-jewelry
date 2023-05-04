package products

import (
	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
)

func (productCategoryUsecase *Usecase) GetAllProductCategories() (*[]ep.ProductCategory, error) {
	productCategories, err := productCategoryUsecase.Repository.GetAllProductCategories()
	return productCategories, err
}

func (productCategoryUsecase *Usecase) GetProductCategoryById(id int) (*ep.ProductCategory, error) {
	productCategory, err := productCategoryUsecase.Repository.GetProductCategoryById(id)
	return productCategory, err
}

func (productCategoryUsecase *Usecase) CreateProductCategory(productCategory *ep.ProductCategory) error {
	err := productCategoryUsecase.Repository.CreateProductCategory(productCategory)
	return err
}

func (productCategoryUsecase *Usecase) UpdateProductCategory(id int, productCategory *ep.ProductCategory) error {
	err := productCategoryUsecase.Repository.UpdateProductCategory(id, productCategory)
	return err
}

func (productCategoryUsecase *Usecase) DeleteProductCategory(id int) error {
	err := productCategoryUsecase.Repository.DeleteProductCategory(id)
	return err
}

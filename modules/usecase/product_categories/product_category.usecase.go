package product_categories

import (
	epc "github.com/berrylradianh/go-jewelry/modules/entity/product_categories"
	rpc "github.com/berrylradianh/go-jewelry/modules/repository/product_categories"
)

type Usecase struct {
	Repository rpc.Repository
}

func (productCategoryUsecase *Usecase) GetAllProductCategories() (*[]epc.ProductCategory, error) {
	productCategories, err := productCategoryUsecase.Repository.GetAllProductCategories()
	return productCategories, err
}

func (productCategoryUsecase *Usecase) GetProductCategoryById(id int) (*epc.ProductCategory, error) {
	productCategory, err := productCategoryUsecase.Repository.GetProductCategoryById(id)
	return productCategory, err
}

func (productCategoryUsecase *Usecase) CreateProductCategory(productCategory *epc.ProductCategory) error {
	err := productCategoryUsecase.Repository.CreateProductCategory(productCategory)
	return err
}

func (productCategoryUsecase *Usecase) UpdateProductCategory(id int, productCategory *epc.ProductCategory) int64 {
	result := productCategoryUsecase.Repository.UpdateProductCategory(id, productCategory)
	return result
}

func (productCategoryUsecase *Usecase) DeleteProductCategory(id int) error {
	err := productCategoryUsecase.Repository.DeleteProductCategory(id)
	return err
}

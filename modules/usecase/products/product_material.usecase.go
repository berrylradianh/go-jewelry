package products

import (
	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
)

func (productMaterialUsecase *Usecase) GetAllProductMaterials() (*[]ep.ProductMaterial, error) {
	productMaterials, err := productMaterialUsecase.Repository.GetAllProductMaterials()
	return productMaterials, err
}

func (productMaterialUsecase *Usecase) GetProductMaterialById(id int) (*ep.ProductMaterial, error) {
	productMaterial, err := productMaterialUsecase.Repository.GetProductMaterialById(id)
	return productMaterial, err
}

func (productMaterialUsecase *Usecase) CreateProductMaterial(productMaterial *ep.ProductMaterial) error {
	err := productMaterialUsecase.Repository.CreateProductMaterial(productMaterial)
	return err
}

func (productMaterialUsecase *Usecase) UpdateProductMaterial(id int, productMaterial *ep.ProductMaterial) error {
	err := productMaterialUsecase.Repository.UpdateProductMaterial(id, productMaterial)
	return err
}

func (productMaterialUsecase *Usecase) DeleteProductMaterial(id int) error {
	err := productMaterialUsecase.Repository.DeleteProductMaterial(id)
	return err
}

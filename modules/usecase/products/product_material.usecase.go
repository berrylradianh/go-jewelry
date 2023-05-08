package products

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (productMaterialUsecase *Usecase) GetAllProductMaterials() (*[]e.ProductMaterial, error) {
	productMaterials, err := productMaterialUsecase.Repository.GetAllProductMaterials()
	return productMaterials, err
}

func (productMaterialUsecase *Usecase) GetProductMaterialById(id int) (*e.ProductMaterial, error) {
	productMaterial, err := productMaterialUsecase.Repository.GetProductMaterialById(id)
	return productMaterial, err
}

func (productMaterialUsecase *Usecase) CreateProductMaterial(productMaterial *e.ProductMaterial) error {
	err := productMaterialUsecase.Repository.CreateProductMaterial(productMaterial)
	return err
}

func (productMaterialUsecase *Usecase) UpdateProductMaterial(id int, productMaterial *e.ProductMaterial) error {
	err := productMaterialUsecase.Repository.UpdateProductMaterial(id, productMaterial)
	return err
}

func (productMaterialUsecase *Usecase) DeleteProductMaterial(id int) error {
	err := productMaterialUsecase.Repository.DeleteProductMaterial(id)
	return err
}

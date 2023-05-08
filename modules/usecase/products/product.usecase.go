package products

import (
	"fmt"

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

func (productUsecase *Usecase) SortProducts(sortBy string, sortOrder string) (*[]e.Product, error) {
	switch sortBy {
	case "name":
		if sortOrder == "asc" {
			products, err := productUsecase.Repository.SortProductByNameASC()
			if err != nil {
				return nil, err
			}
			return products, nil
		} else if sortOrder == "desc" {
			products, err := productUsecase.Repository.SortProductByNameDESC()
			if err != nil {
				return nil, err
			}
			return products, nil
		}
	case "created_at":
		if sortOrder == "asc" {
			products, err := productUsecase.Repository.SortProductByDateASC()
			if err != nil {
				return nil, err
			}
			return products, nil
		} else if sortOrder == "desc" {
			products, err := productUsecase.Repository.SortProductByDateDESC()
			if err != nil {
				return nil, err
			}
			return products, nil
		}
	}
	return nil, fmt.Errorf("invalid field for sorting or invalid sort order")
}

func (productUsecase *Usecase) FilterProductsByMaterial(productMaterial string) (*[]e.Product, error) {
	products, err := productUsecase.Repository.FilterProductsByMaterial(productMaterial)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productUsecase *Usecase) FilterProductsByCategory(productCategory string) (*[]e.Product, error) {
	products, err := productUsecase.Repository.FilterProductsByCategory(productCategory)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productUsecase *Usecase) SearchProductsByName(productName string) ([]e.Product, error) {
	products, err := productUsecase.Repository.SearchProductsByName(productName)
	if err != nil {
		return nil, err
	}
	return products, nil
}

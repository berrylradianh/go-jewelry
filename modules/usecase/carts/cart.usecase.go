package carts

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
	rc "github.com/berrylradianh/go-jewelry/modules/repository/carts"
)

type Usecase struct {
	Repository rc.Repository
}

func (cartUsecase *Usecase) GetAllCarts() (*[]e.Cart, error) {
	cart, err := cartUsecase.Repository.GetAllCarts()
	return cart, err
}

func (cartUsecase *Usecase) GetCartById(id int) (*e.Cart, error) {
	cart, err := cartUsecase.Repository.GetCartById(id)
	return cart, err
}

func (cartUsecase *Usecase) CreateCart(cart *e.Cart) error {
	err := cartUsecase.Repository.CreateCart(cart)
	return err
}

func (cartUsecase *Usecase) UpdateCart(id int, cart *e.Cart) error {
	err := cartUsecase.Repository.UpdateCart(id, cart)
	return err
}

func (cartUsecase *Usecase) DeleteCart(id int) error {
	err := cartUsecase.Repository.DeleteCart(id)
	return err
}

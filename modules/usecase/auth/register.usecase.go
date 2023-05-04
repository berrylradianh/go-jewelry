package auth

import (
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	svc "github.com/berrylradianh/go-jewelry/modules/services"
)

func (authUsecase *Usecase) RegisterUser(user *eu.User) error {
	hashedPassword, err := svc.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = authUsecase.Repository.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}

package auth

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
	svc "github.com/berrylradianh/go-jewelry/modules/services"
)

func (authUsecase *Usecase) RegisterUser(user *e.User) error {
	user.Role_id = 2
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

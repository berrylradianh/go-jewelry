package auth

import (
	"errors"

	"github.com/berrylradianh/go-jewelry/middlewares"
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	svc "github.com/berrylradianh/go-jewelry/modules/services"
)

func (authUsecase *Usecase) LoginUser(email, password string) (*eu.User, string, error) {
	user, err := authUsecase.Repository.LoginUser(email)
	if err != nil {
		return nil, "", err
	}

	err = svc.VerifyPassword(user.Password, password)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, "", errors.New("Invalid Email or Password")
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

package auth

import (
	"github.com/berrylradianh/go-jewelry/middlewares"
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	ra "github.com/berrylradianh/go-jewelry/modules/repository/auth"
	svc "github.com/berrylradianh/go-jewelry/modules/services"
)

type Usecase struct {
	Repository ra.Repository
}

func (authUsecase *Usecase) LoginUser(email, password string) (*eu.User, string, error) {
	user, err := authUsecase.Repository.LoginUser(email)
	if err != nil {
		return nil, "", err
	}

	err = svc.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, "", err
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

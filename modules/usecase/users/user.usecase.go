package users

import (
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	svc "github.com/berrylradianh/go-jewelry/modules/services"
)

func (userUsecase *Usecase) GetAllUsers() (*[]eu.User, error) {
	users, err := userUsecase.Repository.GetAllUsers()
	return users, err
}

func (userUsecase *Usecase) GetUserById(id int) (*eu.User, error) {
	user, err := userUsecase.Repository.GetUserById(id)
	return user, err
}

func (userUsecase *Usecase) CreateUser(user *eu.User) error {
	hashedPassword, err := svc.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = userUsecase.Repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (userUsecase *Usecase) UpdateUser(id int, user *eu.User) error {
	result := userUsecase.Repository.UpdateUser(id, user)
	return result
}

func (userUsecase *Usecase) DeleteUser(id int) error {
	err := userUsecase.Repository.DeleteUser(id)
	return err
}

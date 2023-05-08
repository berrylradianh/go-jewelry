package users

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
	svc "github.com/berrylradianh/go-jewelry/modules/services"
)

func (userUsecase *Usecase) GetAllUsers() (*[]e.User, error) {
	users, err := userUsecase.Repository.GetAllUsers()
	return users, err
}

func (userUsecase *Usecase) GetUserById(id int) (*e.User, error) {
	user, err := userUsecase.Repository.GetUserById(id)
	return user, err
}

func (userUsecase *Usecase) CreateUser(user *e.User) error {
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

func (userUsecase *Usecase) UpdateUser(id int, user *e.User) error {
	result := userUsecase.Repository.UpdateUser(id, user)
	return result
}

func (userUsecase *Usecase) DeleteUser(id int) error {
	err := userUsecase.Repository.DeleteUser(id)
	return err
}

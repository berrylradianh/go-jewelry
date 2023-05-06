package users

import (
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
)

func (userUsecase *Usecase) GetAllUserDetails() (*[]eu.UserDetail, error) {
	userDetails, err := userUsecase.Repository.GetAllUserDetails()
	return userDetails, err
}

func (userUsecase *Usecase) GetUserDetailById(id int) (*eu.UserDetail, error) {
	userDetail, err := userUsecase.Repository.GetUserDetailById(id)
	return userDetail, err
}

func (userUsecase *Usecase) CreateUserDetail(userDetail *eu.UserDetail) error {
	err := userUsecase.Repository.CreateUserDetail(userDetail)
	return err
}

func (userUsecase *Usecase) UpdateUserDetail(id int, userDetail *eu.UserDetail) error {
	result := userUsecase.Repository.UpdateUserDetail(id, userDetail)
	return result
}

func (userUsecase *Usecase) DeleteUserDetail(id int) error {
	err := userUsecase.Repository.DeleteUserDetail(id)
	return err
}

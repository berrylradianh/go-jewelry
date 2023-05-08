package users

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (userUsecase *Usecase) GetAllUserDetails() (*[]e.UserDetail, error) {
	userDetails, err := userUsecase.Repository.GetAllUserDetails()
	return userDetails, err
}

func (userUsecase *Usecase) GetUserDetailById(id int) (*e.UserDetail, error) {
	userDetail, err := userUsecase.Repository.GetUserDetailById(id)
	return userDetail, err
}

func (userUsecase *Usecase) CreateUserDetail(userDetail *e.UserDetail) error {
	err := userUsecase.Repository.CreateUserDetail(userDetail)
	return err
}

func (userUsecase *Usecase) UpdateUserDetail(id int, userDetail *e.UserDetail) error {
	result := userUsecase.Repository.UpdateUserDetail(id, userDetail)
	return result
}

func (userUsecase *Usecase) DeleteUserDetail(id int) error {
	err := userUsecase.Repository.DeleteUserDetail(id)
	return err
}

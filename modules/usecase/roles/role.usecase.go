package roles

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
	rr "github.com/berrylradianh/go-jewelry/modules/repository/roles"
)

type Usecase struct {
	Repository rr.Repository
}

func (roleUsecase *Usecase) GetAllRoles() (*[]e.Role, error) {
	roles, err := roleUsecase.Repository.GetAllRoles()
	return roles, err
}

func (roleUsecase *Usecase) GetRoleById(id int) (*e.Role, error) {
	role, err := roleUsecase.Repository.GetRoleById(id)
	return role, err
}

func (roleUsecase *Usecase) CreateRole(role *e.Role) error {
	err := roleUsecase.Repository.CreateRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (roleUsecase *Usecase) UpdateRole(id int, role *e.Role) error {
	result := roleUsecase.Repository.UpdateRole(id, role)
	return result
}

func (roleUsecase *Usecase) DeleteRole(id int) error {
	err := roleUsecase.Repository.DeleteRole(id)
	return err
}

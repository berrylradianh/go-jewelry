package roles

import (
	// er "github.com/berrylradianh/go-jewelry/modules/entity/roles"
	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	rr "github.com/berrylradianh/go-jewelry/modules/repository/roles"
)

type Usecase struct {
	Repository rr.Repository
}

func (roleUsecase *Usecase) GetAllRoles() (*[]eu.Role, error) {
	roles, err := roleUsecase.Repository.GetAllRoles()
	return roles, err
}

func (roleUsecase *Usecase) GetRoleById(id int) (*eu.Role, error) {
	role, err := roleUsecase.Repository.GetRoleById(id)
	return role, err
}

func (roleUsecase *Usecase) CreateRole(role *eu.Role) error {
	err := roleUsecase.Repository.CreateRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (roleUsecase *Usecase) UpdateRole(id int, role *eu.Role) error {
	result := roleUsecase.Repository.UpdateRole(id, role)
	return result
}

func (roleUsecase *Usecase) DeleteRole(id int) error {
	err := roleUsecase.Repository.DeleteRole(id)
	return err
}

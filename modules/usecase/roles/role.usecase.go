package roles

import (
	er "github.com/berrylradianh/go-jewelry/modules/entity/roles"
	rr "github.com/berrylradianh/go-jewelry/modules/repository/roles"
)

type Usecase struct {
	Repository rr.Repository
}

func (roleUsecase *Usecase) GetAllRoles() (*[]er.Role, error) {
	roles, err := roleUsecase.Repository.GetAllRoles()
	return roles, err
}

func (roleUsecase *Usecase) GetRoleById(id int) (*er.Role, error) {
	role, err := roleUsecase.Repository.GetRoleById(id)
	return role, err
}

func (roleUsecase *Usecase) CreateRole(role *er.Role) error {
	err := roleUsecase.Repository.CreateRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (roleUsecase *Usecase) UpdateRole(id int, role *er.Role) error {
	result := roleUsecase.Repository.UpdateRole(id, role)
	return result
}

func (roleUsecase *Usecase) DeleteRole(id int) error {
	err := roleUsecase.Repository.DeleteRole(id)
	return err
}

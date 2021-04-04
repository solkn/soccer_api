package services

import (
	"github.com/solkn/soccer/api/entity"
	"github.com/solkn/soccer/api/user"
)

type RoleService struct {
	roleRepo user.RoleRepository
}

func NewRoleService(RoleRepo user.RoleRepository) *RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

func (rs *RoleService) Roles() ([]entity.Role, []error) {

	rls, errs := rs.roleRepo.Roles()
	if len(errs) > 0 {
		return nil, errs
	}
	return rls, errs

}

func (rs *RoleService) RoleByName(name string) (*entity.Role, []error) {
	role, errs := rs.roleRepo.RoleByName(name)
	if len(errs) > 0 {
		return nil, errs
	}
	return role, errs
}

func (rs *RoleService) Role(id uint) (*entity.Role, []error) {
	rl, errs := rs.roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

func (rs *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rl, errs := rs.roleRepo.UpdateRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

func (rs *RoleService) DeleteRole(id uint) (*entity.Role, []error) {

	rl, errs := rs.roleRepo.DeleteRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

func (rs *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {

	rl, errs := rs.roleRepo.StoreRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

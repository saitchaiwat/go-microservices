package service

import (
	"user-service/dto"
	"user-service/helper"
	"user-service/model"
	"user-service/repository"
)

type RoleService interface {
	GetAllRole(size int, page int) (paginator helper.Paginator, err error)
	GetRoleByID(id int) (getRole dto.GetRole, err error)
	CreateRole(createRole *dto.CreateRole) error
	UpdateRole(id int, updateRole *dto.UpdateRole) error
	DeleteRole(id int) error
}

type RoleServiceImpl struct {
	roleRepository repository.RoleRepository
}

func (service RoleServiceImpl) GetAllRole(size int, page int) (paginator helper.Paginator, err error) {
	paginator, err = service.roleRepository.GetAllRole(size, page)
	return paginator, err
}

func (service RoleServiceImpl) GetRoleByID(id int) (getRole dto.GetRole, err error) {
	getRole, err = service.roleRepository.GetRoleById(id)
	return getRole, err
}

func (service RoleServiceImpl) CreateRole(createRole *dto.CreateRole) error {
	role := model.Role{Name: createRole.Name}
	return service.roleRepository.CreateRole(role)
}

func (service RoleServiceImpl) UpdateRole(id int, updateRole *dto.UpdateRole) error {
	role := model.Role{Name: updateRole.Name}
	return service.roleRepository.UpdateRole(id, role)
}

func (service RoleServiceImpl) DeleteRole(id int) error {
	return service.roleRepository.DeleteRole(id)
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return RoleServiceImpl{roleRepository: roleRepository}
}

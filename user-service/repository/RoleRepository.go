package repository

import (
	"gorm.io/gorm"
	"user-service/dto"
	"user-service/helper"
	"user-service/model"
)

type RoleRepository interface {
	GetAllRole(size int, page int) (paginator helper.Paginator, err error)
	GetRoleById(id int) (getRole dto.GetRole, err error)
	CreateRole(role model.Role) error
	UpdateRole(id int, role model.Role) error
	DeleteRole(id int) error
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (repository RoleRepositoryImpl) GetAllRole(size int, page int) (paginator helper.Paginator, err error) {
	var getAllRole []dto.GetAllRole
	countSQL := "SELECT COUNT(Id) FROM roles"
	rawSQL := "SELECT id, name FROM roles"
	named := map[string]interface{}{}
	paginator, err = helper.Paginate(repository.db, countSQL, rawSQL, size, page, named, getAllRole)
	return paginator, err
}

func (repository RoleRepositoryImpl) GetRoleById(id int) (getRole dto.GetRole, err error) {
	if err := repository.db.Raw("SELECT * FROM roles WHERE id = ?", id).Scan(&getRole).Error; err != nil {
		return getRole, err
	}
	return getRole, nil
}

func (repository RoleRepositoryImpl) CreateRole(role model.Role) error {
	if err := repository.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (repository RoleRepositoryImpl) UpdateRole(id int, role model.Role) error {
	if err := repository.db.Model(model.Role{}).Where("id = ?", id).Updates(&role).Error; err != nil {
		return err
	}
	return nil
}

func (repository RoleRepositoryImpl) DeleteRole(id int) error {
	if err := repository.db.Delete(model.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return RoleRepositoryImpl{db: db}
}

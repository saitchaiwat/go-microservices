package repository

import (
	"gorm.io/gorm"
	"user-service/dto"
	"user-service/model"
)

type UserRepository interface {
	GetAllUser() (getAllUser []dto.GetAllUser, err error)
	GetUserById()
	CreateUser(user model.User) error
	UpdateUser()
	DeleteUser()
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (repository UserRepositoryImpl) GetAllUser() (getAllUser []dto.GetAllUser, err error) {
	if err := repository.db.Raw("").Scan(getAllUser).Error; err != nil {
		return nil, err
	}
	return getAllUser, nil
}

func (repository UserRepositoryImpl) GetUserById() {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) CreateUser(user model.User) error {
	if err := repository.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repository UserRepositoryImpl) UpdateUser() {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) DeleteUser() {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepositoryImpl{db: db}
}

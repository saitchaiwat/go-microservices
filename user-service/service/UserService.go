package service

import (
	"user-service/dto"
	"user-service/model"
	"user-service/repository"
)

type UserService interface {
	GetAllUser() (getAllUser []dto.GetAllUser, error error)
	GetUserByID()
	CreateUser(createUser *dto.CreateUser) error
	UpdateUser()
	DeleteUser()
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (service UserServiceImpl) GetAllUser() (getAllUser []dto.GetAllUser, error error) {
	return service.userRepository.GetAllUser()
}

func (service UserServiceImpl) GetUserByID() {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) CreateUser(createUser *dto.CreateUser) error {
	user := model.User{
		Username:  createUser.Username,
		Password:  createUser.Password,
		FirstName: createUser.FirstName,
		LastName:  createUser.LastName,
		RoleId:    createUser.RoleId,
	}
	return service.userRepository.CreateUser(user)
}

func (service UserServiceImpl) UpdateUser() {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) DeleteUser() {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserServiceImpl{userRepository: userRepository}
}

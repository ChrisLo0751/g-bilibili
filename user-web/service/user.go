package service

import (
	"g-shopping/user-web/data"
	"g-shopping/user-web/dto"
	"g-shopping/user-web/model"
)

type UserService interface {
	GetUser(id int) (*dto.UserResponse, error)
	CreateUser(userRequest *dto.CreateUserRequest) error
}

type UserServiceImpl struct {
	UserData data.UserData
}

func NewUserService(userData data.UserData) UserService {
	return &UserServiceImpl{UserData: userData}
}

func (u UserServiceImpl) GetUser(id int) (*dto.UserResponse, error) {
	user, err := u.UserData.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		Name: user.Name,
	}, nil
}

func (u UserServiceImpl) CreateUser(userRequest *dto.CreateUserRequest) error {
	return u.UserData.CreateUser(&model.User{
		Name: userRequest.Name,
	})
}

package service

import (
	"g-shopping/user-web/data"
	"g-shopping/user-web/dto"
	"g-shopping/user-web/model"
)

type UserService interface {
	GetUser(id int) (*dto.UserResponse, error)
	CreateUser(user *model.User) error
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

func (u UserServiceImpl) CreateUser(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

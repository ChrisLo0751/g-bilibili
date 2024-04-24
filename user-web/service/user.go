package service

import (
	"fmt"
	"g-shopping/user-web/data"
	"g-shopping/user-web/dto"
	"g-shopping/user-web/model"
)

type UserService interface {
	GetUser(id int) (*model.User, error)
	CreateUser(user *dto.CreateUserRequest) error
}

type UserServiceImpl struct {
	UserData data.UserData
}

func NewUserService(userData data.UserData) UserService {
	return &UserServiceImpl{UserData: userData}
}

func (u UserServiceImpl) GetUser(id int) (*model.User, error) {
	user, err := u.UserData.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    user.ID,
		Phone: user.Phone,
		Email: user.Email,
	}, nil
}

func (u UserServiceImpl) CreateUser(userRequest *dto.CreateUserRequest) error {
	if userRequest.Phone == "" || userRequest.Password == "" {
		return fmt.Errorf("the phone or password is required")
	}

	return u.UserData.CreateUser(&model.User{
		Phone:    userRequest.Phone,
		Password: userRequest.Password,
	})
}

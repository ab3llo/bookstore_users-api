package services

import (
	"github.com/ab3llo/bookstore_users-api/domain/users"
	"github.com/ab3llo/bookstore_users-api/utils/crypto_utils"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	GetAllUsers() ([]*users.User, *errors.RestError)
	UpdateUser(users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) (*users.User, *errors.RestError)
}

// CreateUser creates a user
func (service *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	user.Password = crypto_utils.HashPassword(user.Password)
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser updates user
func (service *usersService) UpdateUser(user users.User) (*users.User, *errors.RestError) {
	user.Password = crypto_utils.HashPassword(user.Password)
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Update(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser using the user Id supplied
func (service *usersService) GetUser(id int64) (*users.User, *errors.RestError) {
	user := users.User{ID: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetAllUsers get users
func (service *usersService) GetAllUsers() ([]*users.User, *errors.RestError) {
	user := users.User{}
	users, err := user.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

//DeleteUser using the user Id supplied
func (service *usersService) DeleteUser(id int64) (*users.User, *errors.RestError) {
	user := users.User{ID: id}
	if err := user.Delete(); err != nil {
		return nil, err
	}
	return &user, nil
}

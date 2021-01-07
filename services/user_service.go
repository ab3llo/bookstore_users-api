package services

import (
	"log"

	"github.com/ab3llo/bookstore_users-api/domain/users"
	"github.com/ab3llo/bookstore_users-api/utils/crypto_utils"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

var (
	//UsersService interface
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(*users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	Login(*users.LoginRequest) (*users.User, *errors.RestError)
	GetAllUsers() ([]*users.User, *errors.RestError)
	UpdateUser(*users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) (*users.User, *errors.RestError)
}

// CreateUser creates a user
func (service *usersService) CreateUser(user *users.User) (*users.User, *errors.RestError) {
	hashedPassword, err := crypto_utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	user.Password = hashedPassword
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	//delete password from response
	user.Password = ""
	return user, nil
}

//UpdateUser updates user
func (service *usersService) UpdateUser(user *users.User) (*users.User, *errors.RestError) {
	hashedPassword, err := crypto_utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	log.Println("Password: " + user.Password + "HashedPassword: " + hashedPassword)
	user.Password = hashedPassword
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Update(); err != nil {
		return nil, err
	}
	//delete password from response
	user.Password = ""
	return user, nil
}

//GetUser using the user Id supplied
func (service *usersService) GetUser(id int64) (*users.User, *errors.RestError) {
	user := &users.User{ID: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

//Login
func (service *usersService) Login(request *users.LoginRequest) (*users.User, *errors.RestError) {
	user := &users.User{Email: request.Email}
	if err := user.FindByEmail(); err != nil {
		return nil, err
	}
	if !crypto_utils.Compare(user.Password, request.Password) {
		// TODO: Properly handle error
		log.Println("Auth failed")
	}
	//delete password from response
	user.Password = ""
	return user, nil
}

//GetAllUsers get users
func (service *usersService) GetAllUsers() ([]*users.User, *errors.RestError) {
	user := &users.User{}
	users, err := user.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

//DeleteUser using the user Id supplied
func (service *usersService) DeleteUser(id int64) (*users.User, *errors.RestError) {
	user := &users.User{ID: id}
	if err := user.Delete(); err != nil {
		return nil, err
	}
	return user, nil
}

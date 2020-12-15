package services

import (
	"github.com/ab3llo/bookstore_users-api/domain/users"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

// CreateUser creates a user
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	//user.CreatedAt = date.GetNow()
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser updates user
func UpdateUser(user users.User) (*users.User, *errors.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Update(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser using the user Id supplied
func GetUser(id int64) (*users.User, *errors.RestError) {
	user := users.User{ID: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

//DeleteUser using the user Id supplied
func DeleteUser(id int64) (*users.User, *errors.RestError) {
	user := users.User{ID: id}
	if err := user.Delete(); err != nil {
		return nil, err
	}
	return &user, nil
}

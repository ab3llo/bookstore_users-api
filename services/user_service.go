package services

import (
	"github.com/ab3llo/bookstore_users-api/domain/users"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

// CreateUser creates a user
func CreateUser (user users.User)(*users.User, *errors.RestError){
  if err := user.Validate(); err != nil {
    return nil, err
  }
  if err := user.Save(); err != nil {
    return nil, err
  }
  return &user, nil
}
package services

import (
	"github.com/ab3llo/bookstore_users-api/domain/users"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

func CreateUser (user users.User)(*users.User, *errors.RestError){

  return &user, nil
}
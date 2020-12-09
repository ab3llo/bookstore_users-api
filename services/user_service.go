package services

import (
	"github.com/ab3llo/bookstore_users-api/domain/users"
	"github.com/ab3llo/bookstore_users-api/utils/date"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

// CreateUser creates a user
func CreateUser (user users.User)(*users.User, *errors.RestError){
  user.DateCreated = date.GetNow()
  if err := user.Validate(); err != nil {
    return nil, err
  }
  if err := user.Save(); err != nil {
    return nil, err
  }
  return &user, nil
}

//GetUser using the user Id supplied 
func GetUser(id int64)(*users.User, *errors.RestError){
  result := users.User{ID: id}
  if err := result.Get(); err !=  nil {
    return nil, err
  }
  return  &result, nil
}
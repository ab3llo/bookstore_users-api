package users

import (
	"fmt"

	database "github.com/ab3llo/bookstore_users-api/datasources/mysql/client"
	"github.com/ab3llo/bookstore_users-api/utils/date"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)


var (
  usersDB = make(map[int64]*User)
)

// Get user by id from db
func (user *User) Get() *errors.RestError {
  if err := database.Client.Ping(); err != nil {
    panic(err)
  }
  result := usersDB[user.ID]
  if result == nil {
    return errors.NewNotFoundError(fmt.Sprintf("User with id: %d not found", user.ID))
  }
  user.ID = result.ID
  user.FirstName = result.FirstName
  user.LastName = result.LastName
  user.Email = result.Email
  user.DateCreated = result.DateCreated
  return nil
}


//Save a user to db
func (user *User) Save() *errors.RestError {
  current := usersDB[user.ID]
  if current != nil {
    if current.Email == user.Email {
      return errors.NewBadRequestError(fmt.Sprintf("email address %s is already registered", user.Email))
    }
    return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
  }
  user.DateCreated = date.GetNow()
  usersDB[user.ID] = user
  return nil
}
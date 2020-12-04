package users

import (
	"strings"

	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

type User struct {
  ID int64 `json:"id"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
  Email string `json:"email"`
  DateCreated string `json:"dateCreated"`
}

// Validate user struct
func(user *User) Validate() *errors.RestError {
  user.Email = strings.TrimSpace(strings.ToLower(user.Email))
  if user.Email == "" {
    return errors.NewBadRequestError("invalid email address")
  }
  return nil
}
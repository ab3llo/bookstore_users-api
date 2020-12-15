package users

import (
	"fmt"

	database "github.com/ab3llo/bookstore_users-api/datasources/mysql/client"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(firstName, lastName, email, dateCreated) VALUES(?,?,?,?);"
)

// Get user by id from db
func (user *User) Get() *errors.RestError {
	result := database.Client.First(user, user.ID)
	if result.Error != nil {
		return errors.NewNotFoundError(fmt.Sprintf("User with id: %d not found", user.ID))
	}
	return nil
}

//Save a user to db
func (user *User) Save() *errors.RestError {
	result := database.Client.Create(user)
	if result.Error != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", result.Error.Error()))
	}
	return nil
}

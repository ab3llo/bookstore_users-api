package users

import (
	"fmt"
	"log"

	database "github.com/ab3llo/bookstore_users-api/datasources/mysql/client"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

// Get user by id from db
func (user *User) Get() *errors.RestError {
	result := database.Client.First(user, user.ID)
	if result.Error != nil {
		return errors.NewNotFoundError(fmt.Sprintf("User with id: %d not found", user.ID))
	}
	return nil
}

//GetAll users
func (user *User) GetAll() ([]*User, *errors.RestError) {
	var records []*User
	if db := database.Client.Find(&records); db.Error != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Error %s", db.Error))
	}
	return records, nil
}

//Save a user to db
func (user *User) Save() *errors.RestError {
	result := database.Client.Create(user)
	if result.Error != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", result.Error.Error()))
	}
	log.Printf("Created user record with id: %d", user.ID)
	return nil
}

//Update a user in the db
func (user *User) Update() *errors.RestError {
	var currentUser = &User{ID: user.ID}
	result := database.Client.First(currentUser)
	if result.Error != nil {
		return errors.NewNotFoundError(
			fmt.Sprintf("User with id: %d not found", user.ID))
	}
	user.CreatedAt = currentUser.CreatedAt
	err := database.Client.Model(&user).Updates(user).Error
	if err != nil {
		return errors.NewBadRequestError(
			fmt.Sprintf("error when trying to update user: %s", err.Error()))
	}

	log.Printf("Updated user record with id: %d", user.ID)
	return nil
}

//Delete User in db
func (user *User) Delete() *errors.RestError {
	result := database.Client.Unscoped().Delete(user)
	if result.Error != nil {
		return errors.NewNotFoundError(
			fmt.Sprintf("User with id: %d not found", user.ID))
	}
	log.Printf("Deleted user record with id: %d", user.ID)
	return nil
}

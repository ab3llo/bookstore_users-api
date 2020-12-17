package users

import (
	"log"

	"github.com/ab3llo/bookstore_users-api/utils/mysql_utils"

	database "github.com/ab3llo/bookstore_users-api/datasources/mysql/client"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

// Get user by id from db
func (user *User) Get() *errors.RestError {
	result := database.Client.First(user, user.ID)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

//GetAll users
func (user *User) GetAll() ([]*User, *errors.RestError) {
	var records []*User
	if db := database.Client.Find(&records); db.Error != nil {
		return nil, mysql_utils.ParseError(db.Error)
	}
	return records, nil
}

//Save a user to db
func (user *User) Save() *errors.RestError {
	result := database.Client.Create(user)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	log.Printf("Created user record with id: %d", user.ID)
	return nil
}

//Update a user in the db
func (user *User) Update() *errors.RestError {
	var currentUser = &User{ID: user.ID}
	result := database.Client.First(currentUser)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	user.CreatedAt = currentUser.CreatedAt
	err := database.Client.Model(&user).Updates(user).Error
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	log.Printf("Updated user record with id: %d", user.ID)
	return nil
}

//Delete User in db
func (user *User) Delete() *errors.RestError {
	result := database.Client.Unscoped().Delete(user)
	if result.Error != nil {
		return mysql_utils.ParseError(result.Error)
	}
	log.Printf("Deleted user record with id: %d", user.ID)
	return nil
}

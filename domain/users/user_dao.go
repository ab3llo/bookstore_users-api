package users

import (
	"github.com/ab3llo/bookstore_users-api/logger"
	"github.com/ab3llo/bookstore_users-api/utils/mysql_utils"

	database "github.com/ab3llo/bookstore_users-api/datasources/mysql/client"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

// Get user by id from db
func (user *User) Get() *errors.RestError {
	result := database.Client.First(user, user.ID)
	if result.Error != nil {
		logger.Error("error when trying to get user details", result.Error)
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

//FindByEmail find user by email
func (user *User) FindByEmail() *errors.RestError {
	result := database.Client.Where("email =? AND status=?", user.Email, "active").First(&user)
	if result.Error != nil {
		logger.Error("error when trying to find user by email", result.Error)
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

//GetAll users
func (user *User) GetAll() ([]*User, *errors.RestError) {
	var records []*User
	if db := database.Client.Find(&records); db.Error != nil {
		logger.Error("error when trying to get users details", db.Error)
		return nil, mysql_utils.ParseError(db.Error)
	}
	return records, nil
}

//Save a user to db
func (user *User) Save() *errors.RestError {
	user.Status = StatusActive
	result := database.Client.Create(user)
	if result.Error != nil {
		logger.Error("error when trying to get save user details", result.Error)
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

//Update a user in the db
func (user *User) Update() *errors.RestError {
	var currentUser = &User{ID: user.ID}
	result := database.Client.First(currentUser)
	if result.Error != nil {
		logger.Error("error when trying to get user details", result.Error)
		return mysql_utils.ParseError(result.Error)
	}
	user.CreatedAt = currentUser.CreatedAt
	err := database.Client.Save(&user).Error
	if err != nil {
		logger.Error("error when trying to update user details", result.Error)
		return mysql_utils.ParseError(err)
	}
	return nil
}

//Delete User in db
func (user *User) Delete() *errors.RestError {
	var currentUser = &User{ID: user.ID}
	result := database.Client.First(currentUser)
	if result.Error != nil {
		logger.Error("error when trying to get user details", result.Error)
		return mysql_utils.ParseError(result.Error)
	}
	result = database.Client.Unscoped().Delete(user)
	if result.Error != nil {
		logger.Error("error when trying to delete user details", result.Error)
		return mysql_utils.ParseError(result.Error)
	}
	return nil
}

package users

import (
	"fmt"

	database "github.com/ab3llo/bookstore_users-api/datasources/mysql/client"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
)

const (
  queryInsertUser ="INSERT INTO users(firstName, lastName, email, dateCreated) VALUES(?,?,?,?);"
)

// Get user by id from db
func (user *User) Get() *errors.RestError {
  if err := database.Client.Ping(); err != nil {
    panic(err)
  }
  result := user
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
  statement, err := database.Client.Prepare(queryInsertUser)
  if err != nil {
    return errors.NewInternalServerError(err.Error())
  }
  defer statement.Close()

  insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
  if err != nil {
    return errors.NewInternalServerError(
      fmt.Sprintf("error when trying to save user: %s",err.Error()))
  }
  
  userID, err := insertResult.LastInsertId()
  if err != nil {
    return errors.NewInternalServerError(
      fmt.Sprintf("error when trying to save user: %s",err.Error()))
  }
  user.ID = userID
  return nil
}
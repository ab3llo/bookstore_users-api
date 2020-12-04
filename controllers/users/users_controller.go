package users

import (
	"net/http"

	"github.com/ab3llo/bookstore_users-api/domain/users"
	"github.com/ab3llo/bookstore_users-api/services"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//Each controller function needs to implement *gin.Context interface

// CreateUser creates a user
func CreateUser(c *gin.Context){
  var user users.User 
  if err := c.ShouldBindJSON(&user); err != nil {
    restErr := errors.NewBadRequestError("invalid json body")
    c.JSON(restErr.Status, restErr)
    return 
  }

  result, saveErr := services.CreateUser(user)
  if saveErr != nil {
    c.JSON(saveErr.Status, saveErr)
    return
  }
  c.JSON(http.StatusCreated, result)
}

// GetUser gets a user using the id
func GetUser(c *gin.Context){
  
}

package app

import (
	"github.com/ab3llo/bookstore_users-api/controllers/ping"
	"github.com/ab3llo/bookstore_users-api/controllers/users"
)

func mapUrls(){
  router.GET("/_ping", ping.Ping)
  router.POST("/users/", users.CreateUser)
  router.GET("/users/:user_id", users.GetUser)
}
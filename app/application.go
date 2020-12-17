package app

import (
	"github.com/ab3llo/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication starts the application
func StartApplication() {
	mapUrls()
	logger.Info("Starting application....")
	router.Run(":8080")
}

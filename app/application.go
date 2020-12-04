package app

import "github.com/gin-gonic/gin"

var (
  router = gin.Default()
)
// StartApplication starts the application
func StartApplication(){
  mapUrls()
  router.Run(":8080")
}
package main

import (
	"log"

	"github.com/ab3llo/bookstore_users-api/app"
	"github.com/joho/godotenv"
)

func init(){
  if err := godotenv.Load(); err != nil {
    log.Print("No .env file found")
  }
}
func main(){
  app.StartApplication()
}
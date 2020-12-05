package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ab3llo/bookstore_users-api/config"
	// import go sql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
  //Client for database
  Client *sql.DB
)


func init(){
  godotenv.Load()
  conf := config.New()
  dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", 
    conf.Datbase.Username, 
    conf.Datbase.Password,
    conf.Datbase.Host,
    conf.Datbase.Schema,
  )
  Client, err := sql.Open("mysql", dataSourceName)
  if err != nil {
    panic(err)
  }
  if err = Client.Ping(); err != nil {
    panic(err)
  }
  log.Println("database sucessfully configured.")
}
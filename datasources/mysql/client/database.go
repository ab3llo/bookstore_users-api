package database

import (
	"fmt"
	"log"

	"github.com/ab3llo/bookstore_users-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// import go sql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	//Client for database
	Client  *gorm.DB
	dbError error
)

func init() {
	godotenv.Load()
	conf := config.New()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		conf.Datbase.Username,
		conf.Datbase.Password,
		conf.Datbase.Host,
		conf.Datbase.Schema,
	)
	Client, dbError = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if dbError != nil {
		panic(dbError)
	}
	Client.DB()
	log.Println("database successfully configured.")
}

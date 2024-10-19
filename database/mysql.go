package database

import (
	"fmt"
	"test-server/dbmodel"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DBInstance *gorm.DB
var err error
var CONNECTION_STRING string = "@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local"

func ConnectDB(username string, passwd string) {
	ConnectionURI := fmt.Sprintf("%s:%s%s", username, passwd, CONNECTION_STRING)

	DBInstance, err = gorm.Open("mysql", ConnectionURI)
	if err != nil {
		fmt.Println(err)
		panic("Database connection attempt was unsuccessful.....")
	} else {
		fmt.Println("Database Connected successfully.....")
	}

	DBInstance.LogMode(true)
}

func CreateDB() {
	DBInstance.Exec("CREATE DATABASE IF NOT EXISTS swapidata")
	DBInstance.Exec("USE swapidata")
}

func MigrateDB() {
	DBInstance.AutoMigrate(&dbmodel.Vehicle{})
	fmt.Println("Database migration completed....")
}

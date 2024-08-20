package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectToDB() {
	connection, errConnection := gorm.Open("mysql", "myroot:password@tcp(localhost:3306)/ticket_manager?charset=utf8&parseTime=True&loc=Local")

	if errConnection != nil {
		fmt.Println("errore while connection to database  on app.go: ", errConnection)
		panic(errConnection)
	}
	db = connection
	fmt.Printf("Connection to database success...listing ad port : 3306")
}

func GetDB() *gorm.DB {
	return db
}

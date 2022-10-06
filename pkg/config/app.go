package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	fmt.Println("Connecting to database...")
	
	connString := "Adema:YES@tcp(127.0.0.1:3306)/databaseName"
	d, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}

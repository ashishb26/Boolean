package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func connectDb() {
	db, err = gorm.Open("mysql", "ashish:root@tcp(127.0.0.1:3306)/tempdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	db.AutoMigrate(&BoolTable{})
}

package dbConfig

import (
	"log"

	"github.com/Boolean/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func ConnectDb() {
	DB, err = gorm.Open("mysql", "ashish:root@tcp(127.0.0.1:3306)/tempdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	DB.AutoMigrate(&models.BoolTable{})
}

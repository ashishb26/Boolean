package dbConfig

import (
	"log"

	"github.com/Boolean/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/xid"
)

// DB represents the database object to which a connection
// has been established
var DB *gorm.DB
var err error

// ConnectDb connects to the database and creates the boot_table and credentials tables
// (if they are not already present in the mysql database)
func ConnectDb() {
	DB, err = gorm.Open("mysql", "ashish:root@tcp(127.0.0.1:3306)/tempdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	DB.AutoMigrate(&models.BoolTable{})
	DB.AutoMigrate(&models.Credentials{})

	var userCred models.Credentials
	userCred.ID = xid.New().String()
	userCred.Username = "root"
	userCred.Password = "password"

	DB.Create(&userCred)
}

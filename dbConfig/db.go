package dbConfig

import (
	"log"
	"sync"

	"github.com/Boolean/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// DB represents the database to which a connection
// has been established
var DB *gorm.DB
var err error

var dbUserName = "root"
var dbPassword = ""
var dbName = "booldb"

// Mu is a RWMutex used to synchronize database read and writes
var Mu *sync.RWMutex

// ConnectDb connects to the database and creates the bool_table and credentials tables
// (if they are not already present in the mysql database) and stores this connection interface{}
// the variable 'DB'
func ConnectDb() {
	DB, err = gorm.Open("mysql", dbUserName+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")

	Mu = &sync.RWMutex{}

	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	DB.AutoMigrate(&models.BoolTable{})
	DB.AutoMigrate(&models.Credentials{})

	var record models.Credentials

	// Check if (root,password) is already present in the database
	err := DB.Where("username=?", "root").First(&record).Error
	if gorm.IsRecordNotFoundError(err) {

		var userCred models.Credentials
		userCred.ID = xid.New().String()
		userCred.Username = "root"
		userCred.Password = "password"

		DB.Create(&userCred)
		return

	} else if err != nil {
		log.Fatalln(err)
	}

}

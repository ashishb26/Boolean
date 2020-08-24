package controller

import (
	"log"

	"github.com/ashishb26/rzpbool/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
)

// Server struct is used to represent the database connection and router
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// NewServer creates a new database connection and router and returns a Server struct
func NewServer() *Server {

	dbConfig := models.GetDBConfig()

	db, err := gorm.Open("mysql", dbConfig.DBUser+":"+dbConfig.DBPassword+"@tcp("+dbConfig.DBHost+":"+dbConfig.DBPort+")/"+dbConfig.DBName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	router := gin.Default()

	db.AutoMigrate(&models.BoolRecord{})

	return &Server{
		DB:     db,
		Router: router,
	}
}

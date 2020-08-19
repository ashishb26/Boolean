package main

import (
	"github.com/Boolean/controller"
	"github.com/Boolean/dbConfig"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dbConfig.ConnectDb()
	defer dbConfig.DB.Close()
	router.POST("/", controller.AddBool)
	router.GET("/:id", controller.GetBool)
	router.PATCH("/:id", controller.UpdateBool)
	router.DELETE("/:id", controller.DeleteBool)
	router.POST("/login", controller.Login)
	router.POST("/adduser", controller.AddUser)
	router.Run(":8080")
}

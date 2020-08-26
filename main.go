package main

import (
	"github.com/ashishb26/rzpbool/controller"
)

func main() {
	newServer := controller.NewServer()
	newServer.InitRoutes()
	defer newServer.DB.Close()
	newServer.Router.Run(":8080")
}

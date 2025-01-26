package main

import (
	"example.com/REST/db"
	"example.com/REST/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //Running on local host

}

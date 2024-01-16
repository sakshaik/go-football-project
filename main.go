package main

import (
	"example.com/football-project/db"
	"example.com/football-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

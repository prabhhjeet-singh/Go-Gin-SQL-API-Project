package main

import (
	"github.com/gin-gonic/gin"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
)

func main() {

	db.InitDB()
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run("localhost:8081")
}

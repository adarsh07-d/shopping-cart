package main

import (
	"github.com/gin-gonic/gin"
	"shopping-cart-backend/controllers"
	"shopping-cart-backend/database"
)

func main() {
	database.Connect()

	r := gin.Default()

	// Register routes directly
	r.POST("/users", controllers.CreateUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/items", controllers.ListItems)

	// Run the server
	r.Run(":9090")
}

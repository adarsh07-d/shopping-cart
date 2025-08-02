package routes

import (
	"github.com/gin-gonic/gin"
	"shopping-cart-backend/controllers"
	"shopping-cart-backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// 🧑‍💻 Public Routes
	r.POST("/users", controllers.CreateUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/users", controllers.ListUsers)

	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.ListItems)

	// 🔒 Protected Routes (require token)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/carts", controllers.AddToCart)
		auth.GET("/carts", controllers.ListCarts)
		auth.POST("/orders", controllers.CreateOrder)
		auth.GET("/orders", controllers.ListOrders)
	}
}

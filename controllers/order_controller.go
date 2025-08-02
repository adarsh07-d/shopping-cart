package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart-backend/database"
	"shopping-cart-backend/models"
)

func CreateOrder(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var cart models.Cart
	if err := database.DB.Preload("Items").Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	order := models.Order{
		UserID: user.ID,
		Items:  cart.Items,
	}
	database.DB.Create(&order)

	// Clear cart items after order is placed
	database.DB.Model(&cart).Association("Items").Clear()

	c.JSON(http.StatusOK, gin.H{"message": "Order placed", "order_id": order.ID})
}

func ListOrders(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var orders []models.Order
	database.DB.Preload("Items").Where("user_id = ?", user.ID).Find(&orders)

	c.JSON(http.StatusOK, orders)
}

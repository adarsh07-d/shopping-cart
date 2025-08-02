package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart-backend/database"
	"shopping-cart-backend/models"
)

type AddToCartInput struct {
	ItemID uint `json:"item_id"`
}

func AddToCart(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var input AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var item models.Item
	if err := database.DB.First(&item, input.ItemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var cart models.Cart
	if err := database.DB.Preload("Items").Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		cart = models.Cart{UserID: user.ID}
		database.DB.Create(&cart)
	}

	database.DB.Model(&cart).Association("Items").Append(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func ListCarts(c *gin.Context) {
	var carts []models.Cart
	database.DB.Preload("Items").Find(&carts)
	c.JSON(http.StatusOK, carts)
}

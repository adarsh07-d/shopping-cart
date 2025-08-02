package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart-backend/database"
	"shopping-cart-backend/models"
)

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&item)
	c.JSON(http.StatusCreated, item)
}

func ListItems(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

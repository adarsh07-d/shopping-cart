package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"shopping-cart-backend/models"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("shopping.db"), &gorm.Config{})
	if err != nil {
		panic("❌ Failed to connect to DB")
	}
	db.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})
	DB = db
}

package models

type Order struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User
	Items  []Item `gorm:"many2many:order_items;"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartItemID uint      `gorm:"column:cart_item_id;primaryKey;autoIncrement"`
	CartID     uint      `gorm:"column:cart_id;not null"`
	ProductID  uint      `gorm:"column:product_id;not null"`
	Quantity   int       `gorm:"column:quantity;not null"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	Cart       Cart      `gorm:"foreignKey:CartID"`
	Product    Product   `gorm:"foreignKey:ProductID"`
}

func (CartItem) TableName() string {
	return "cart_item"
}

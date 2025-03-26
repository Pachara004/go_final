package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CartID     uint       `gorm:"column:cart_id;primaryKey;autoIncrement"`
	CustomerID uint       `gorm:"column:customer_id;not null"`
	CartName   string     `gorm:"column:cart_name"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	Customer   Customer   `gorm:"foreignKey:CustomerID"`
	CartItems  []CartItem `gorm:"foreignKey:CartID"`
}

func (Cart) TableName() string {
	return "cart"
}

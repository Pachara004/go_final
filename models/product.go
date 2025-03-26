package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID     uint      `gorm:"column:product_id;primaryKey;autoIncrement"`
	ProductName   string    `gorm:"column:product_name;not null"`
	Description   string    `gorm:"column:description"`
	Price         float64   `gorm:"column:price;not null"`
	StockQuantity int       `gorm:"column:stock_quantity;not null"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "product"
}

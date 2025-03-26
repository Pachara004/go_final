package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerID  uint      `gorm:"column:customer_id;primaryKey;autoIncrement"`
	FirstName   string    `gorm:"column:first_name;not null"`
	LastName    string    `gorm:"column:last_name;not null"`
	Email       string    `gorm:"column:email;uniqueIndex;not null"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Address     string    `gorm:"column:address"`
	Password    string    `gorm:"column:password;not null"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Customer) TableName() string {
	return "customer"
}

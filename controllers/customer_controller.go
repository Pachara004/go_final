package controllers

import (
	"go_final/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerController struct {
	db *gorm.DB
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{db: db}
}

func (cc *CustomerController) GetProfile(c *gin.Context) {
	customerID := c.Param("customer_id")

	var customer models.Customer
	result := cc.db.First(&customer, customerID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"customer": gin.H{
			"customer_id":  customer.CustomerID,
			"email":        customer.Email,
			"first_name":   customer.FirstName,
			"last_name":    customer.LastName,
			"phone_number": customer.PhoneNumber,
			"address":      customer.Address,
		},
	})
}

func (cc *CustomerController) UpdateProfile(c *gin.Context) {
	customerID := c.Param("customer_id")

	var updateData struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	result := cc.db.First(&customer, customerID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	customer.FirstName = updateData.FirstName
	customer.LastName = updateData.LastName
	customer.PhoneNumber = updateData.PhoneNumber
	customer.Address = updateData.Address

	cc.db.Save(&customer)

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"customer": gin.H{
			"first_name":   customer.FirstName,
			"last_name":    customer.LastName,
			"phone_number": customer.PhoneNumber,
			"address":      customer.Address,
		},
	})
}

package controllers

import (
	"go_final/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		log.Println("Invalid JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	result := ac.DB.Where("email = ?", loginRequest.Email).First(&customer)
	if result.Error != nil {
		log.Printf("Query failed for email %s: %v", loginRequest.Email, result.Error)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "อีเมลหรือรหัสผ่านไม่ถูกต้อง"})
		return
	}

	log.Printf("Found customer: %+v", customer)

	c.JSON(http.StatusOK, gin.H{
		"customer_id": customer.CustomerID,
		"first_name":  customer.FirstName,
		"last_name":   customer.LastName,
		"email":       customer.Email,
		"phonenumber": customer.PhoneNumber,
		"address":     customer.Address,
		"createat":    customer.CreatedAt,
		"updateat":    customer.UpdatedAt,
	})
}
func (ac *AuthController) ChangePassword(c *gin.Context) {
	var passwordChangeRequest struct {
		CustomerID  uint   `json:"customer_id" binding:"required"`
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&passwordChangeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	result := ac.DB.First(&customer, passwordChangeRequest.CustomerID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบผู้ใช้"})
		return
	}

	// อัปเดตรหัสผ่านใหม่
	ac.DB.Save(&customer)

	c.JSON(http.StatusOK, gin.H{"message": "เปลี่ยนรหัสผ่านสำเร็จ"})
}

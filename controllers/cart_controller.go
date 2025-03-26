package controllers

import (
	"go_final/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartController struct {
	DB *gorm.DB
}

func NewCartController(db *gorm.DB) *CartController {
	return &CartController{DB: db}
}

func (cc *CartController) AddToCart(c *gin.Context) {
	var cartRequest struct {
		CustomerID uint   `json:"customer_id" binding:"required"`
		CartName   string `json:"cart_name"`
		ProductID  uint   `json:"product_id" binding:"required"`
		Quantity   int    `json:"quantity" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&cartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหาหรือสร้างรถเข็น
	var cart models.Cart
	result := cc.DB.Where("customer_id = ? AND cart_name = ?",
		cartRequest.CustomerID,
		cartRequest.CartName).FirstOrCreate(&cart,
		models.Cart{
			CustomerID: cartRequest.CustomerID,
			CartName:   cartRequest.CartName,
		})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้างรถเข็นได้"})
		return
	}

	// ตรวจสอบสินค้าในรถเข็น
	var existingCartItem models.CartItem
	existingResult := cc.DB.Where("cart_id = ? AND product_id = ?",
		cart.CartID,
		cartRequest.ProductID).First(&existingCartItem)

	if existingResult.Error == nil {
		// อัปเดตจำนวนสินค้าถ้ามีอยู่แล้ว
		existingCartItem.Quantity += cartRequest.Quantity
		cc.DB.Save(&existingCartItem)
	} else {
		// เพิ่มสินค้าใหม่ในรถเข็น
		newCartItem := models.CartItem{
			CartID:    cart.CartID,
			ProductID: cartRequest.ProductID,
			Quantity:  cartRequest.Quantity,
		}
		cc.DB.Create(&newCartItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "เพิ่มสินค้าในรถเข็นสำเร็จ"})
}

func (cc *CartController) ListCarts(c *gin.Context) {
	customerID := c.Param("customer_id")

	var carts []models.Cart
	err := cc.DB.Preload("CartItems.Product").
		Where("customer_id = ?", customerID).
		Find(&carts).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูลรถเข็นได้"})
		return
	}

	c.JSON(http.StatusOK, carts)
}

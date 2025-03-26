package controllers

import (
	"go_final/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) SearchProducts(c *gin.Context) {
	var searchRequest struct {
		Keyword  string  `form:"keyword"`
		MinPrice float64 `form:"min_price"`
		MaxPrice float64 `form:"max_price"`
	}

	if err := c.ShouldBindQuery(&searchRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var products []models.Product
	query := pc.DB.Model(&models.Product{})

	// กรองตามคำค้นหา
	if searchRequest.Keyword != "" {
		query = query.Where("product_name LIKE ? OR description LIKE ?",
			"%"+searchRequest.Keyword+"%",
			"%"+searchRequest.Keyword+"%")
	}

	// กรองตามช่วงราคา
	if searchRequest.MinPrice > 0 {
		query = query.Where("price >= ?", searchRequest.MinPrice)
	}
	if searchRequest.MaxPrice > 0 {
		query = query.Where("price <= ?", searchRequest.MaxPrice)
	}

	// ค้นหาสินค้า
	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ค้นหาสินค้าล้มเหลว"})
		return
	}

	c.JSON(http.StatusOK, products)
}

package main

import (
	"go_final/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("อ่านไฟล์การตั้งค่าผิดพลาด: %v", err)
	}

	dsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("เชื่อมต่อฐานข้อมูลล้มเหลว: %v", err)
	}
	log.Println("เชื่อมต่อฐานข้อมูลสำเร็จ") // เพิ่มเพื่อยืนยัน

	r := gin.Default()
	authController := controllers.NewAuthController(db)
	productController := controllers.NewProductController(db)
	cartController := controllers.NewCartController(db)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authController.Login)
		v1.PUT("/change-password", authController.ChangePassword)
		v1.GET("/products/search", productController.SearchProducts)
		v1.POST("/cart/add", cartController.AddToCart)
		v1.GET("/cart/list/:customer_id", cartController.ListCarts)
	}

	r.Run(":8080")
}

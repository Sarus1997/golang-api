package main

import (
	"github.com/Sarus1997/golang-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	//* สร้าง Gin router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "N O O B",
		})
	})

  // Route สำหรับดึงข้อมูลสินค้าทั้งหมด
	r.GET("/products", handlers.GetProducts)

	// Route สำหรับดึงข้อมูลสินค้าตาม product_id
	r.GET("/products/:product_id", handlers.GetProductByID)

	//* routes สำหรับ authentication
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}

	//* port 8080
	r.Run(":8080")
}


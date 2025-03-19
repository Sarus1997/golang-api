package handlers

import (
	"net/http"

	"github.com/Sarus1997/golang-api/config"
	"github.com/Sarus1997/golang-api/models"

	"github.com/gin-gonic/gin"
)

// ฟังก์ชันสำหรับดึงข้อมูลสินค้า
func GetProducts(c *gin.Context) {
	var products []models.Product

	// ดึงข้อมูลทั้งหมดจากตาราง product_
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}

	// ส่งข้อมูลสินค้าให้กับผู้ใช้
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// ฟังก์ชันสำหรับดึงข้อมูลสินค้าตาม product_id
func GetProductByID(c *gin.Context) {
	productID := c.Param("product_id")

	var product models.Product

	// ดึงข้อมูลสินค้าตาม product_id
	if err := config.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// ส่งข้อมูลสินค้าให้กับผู้ใช้
	c.JSON(http.StatusOK, gin.H{"product": product})
}

package handlers

import (
	"net/http"

	"github.com/Sarus1997/golang-api/config"
	"github.com/Sarus1997/golang-api/models"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	//* ดึงข้อมูลทั้งหมดจากตาราง test_101 = Test
	var test []models.Test
	if err := config.DB.Find(&test).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"test": test})
}

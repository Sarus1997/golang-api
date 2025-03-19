package handlers

import (
	"net/http"

	"github.com/Sarus1997/golang-api/config"
	"github.com/Sarus1997/golang-api/models"

	"github.com/gin-gonic/gin"
)

//* Get
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

// * Post
func PostTest(c *gin.Context) {
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Create(&test).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating test"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Test created successfully"})
}

// * Put
func PutTest(c *gin.Context) {
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var test models.Test
	if err := config.DB.Where("id = ?", id).First(&test).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test not found"})
		return
	}

	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Save(&test).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test updated successfully"})
}

// * Delete
func DeleteTest(c *gin.Context) {
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var test models.Test
	if err := config.DB.Where("id = ?", id).First(&test).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test not found"})
		return
	}

	if err := config.DB.Delete(&test).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test deleted successfully"})
}


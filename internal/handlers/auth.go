package handlers

import (
	"net/http"

	"github.com/Sarus1997/golang-api/config"
	"github.com/Sarus1997/golang-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ฟังก์ชันสำหรับการลงทะเบียนผู้ใช้
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// เข้ารหัสรหัสผ่านก่อนบันทึก
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// สร้างผู้ใช้ในฐานข้อมูล
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// ฟังก์ชันสำหรับการเข้าสู่ระบบ
func Login(c *gin.Context) {
	var loginData models.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// ตรวจสอบข้อมูลผู้ใช้จากฐานข้อมูล
	var user models.User
	if err := config.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// ส่ง JWT Token กลับ
	// token := generateJWT(user.Username) // ฟังก์ชันการสร้าง JWT Token ที่คุณจะสร้างขึ้นเอง
	// c.JSON(http.StatusOK, gin.H{"token": token})

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

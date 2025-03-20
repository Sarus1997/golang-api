package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/Sarus1997/golang-api/config"
	"github.com/Sarus1997/golang-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// * ฟังก์ชันสำหรับการสร้าง ID สุ่ม
func generateID() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

//* ฟังก์ชันสำหรับการลงทะเบียนผู้ใช้
func Register(c *gin.Context) {
  var registerData struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	FirstName      string `json:"f_name"`
	LastName       string `json:"l_name"`
	ProfilePicture string `json:"profile_picture"`
  }

  if err := c.ShouldBindJSON(&registerData); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  //* สร้าง id อัตโนมัต
  id := generateID()

  //* เข้ารหัสรหัสผ่านก่อนบันทึก
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
    return
  }

  user := models.User{
    ID:             id,
    Username:       registerData.Username,
		Email:          registerData.Email,
    PasswordHash:   string(hashedPassword),
    FirstName:      registerData.FirstName,
    LastName:       registerData.LastName,
    ProfilePicture: registerData.ProfilePicture,
  }

  //* สร้างผู้ใช้ในฐานข้อมูล
  if err := config.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// * ฟังก์ชันสำหรับการเข้าสู่ระบบ
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// ตรวจสอบ JSON request
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// ตรวจสอบว่ามีผู้ใช้งานอยู่ในฐานข้อมูลหรือไม่
	var user models.User
	if err := config.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		fmt.Println("User not found:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Debug: ตรวจสอบรหัสผ่านก่อนเปรียบเทียบ
	fmt.Println("Input Password:", loginData.Password)
	fmt.Println("Stored Hash:", user.PasswordHash)

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginData.Password))
	if err != nil {
		fmt.Println("Password does not match:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// อัปเดตเวลาล็อกอินล่าสุด
	now := time.Now()
	config.DB.Model(&user).Update("last_login_at", now)

	// ส่ง JWT Token กลับ (สามารถเพิ่มฟังก์ชัน generateJWT)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

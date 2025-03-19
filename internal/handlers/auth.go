package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/Sarus1997/golang-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Fake DB (ใช้จริงต้องเชื่อม Database)
var users = []models.User{
	{
		Username: "user1",
		Password: "password1",
	},
}
var db *gorm.DB

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	users = append(users, user)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var loginData models.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for _, user := range users {
		if user.Username == loginData.Username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
			if err == nil {
				token := generateJWT(user.Username)
				c.JSON(http.StatusOK, gin.H{"token": token})
				return
			}
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func generateJWT(username string) string {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString
}

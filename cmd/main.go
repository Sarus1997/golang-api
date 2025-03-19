package main

import (
	"github.com/Sarus1997/golang-api/config"
	"github.com/Sarus1997/golang-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}

	r.Run(":8080")
}

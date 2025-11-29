package main

import (
	"github.com/gin-gonic/gin"
	"kotiki/handlers"
	"kotiki/users"
)

func main() {
	r := gin.Default()

	// 1. создаём репозиторий пользователей
	userRepo := users.NewInMemory()

	// 2. создаём AuthService
	authService := users.NewAuthService(userRepo)

	// 3. создаём UserHandlers
	userHandlers := handlers.NewUserHandlers(authService)

	// 4. маршрут для регистрации
	r.POST("/register", userHandlers.Register)

	// (/ping по приколу)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	if err := r.Run(":8080"); err != nil {
		return
	}
}

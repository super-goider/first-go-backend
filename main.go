package main

import (
	"github.com/gin-gonic/gin"

	"kotiki/cats"
	"kotiki/handlers"
	"kotiki/users"
)

func main() {
	r := gin.Default()

	// users

	userRepo := users.NewInMemory()
	authService := users.NewAuthService(userRepo)
	userHandlers := handlers.NewUserHandlers(authService)

	r.POST("/register", userHandlers.Register)
	r.POST("/login", userHandlers.Login)
	r.GET("/me", handlers.AuthRequired, userHandlers.Me)

	// cats

	catRepo := cats.NewInMemory()
	catHandlers := handlers.NewCatHandlers(catRepo)

	// список котов (без авторизации пока)
	r.GET("/cats", catHandlers.GetAllCat)

	// to be done
	// r.POST("/cats", handlers.AuthRequired, catHandlers.CreateCat)
	// r.GET("/cats/:id", catHandlers.GetCat)
	// r.DELETE("/cats/:id", handlers.AuthRequired, catHandlers.DeleteCat)

	// тест
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	if err := r.Run(":8080"); err != nil {
		return
	}
}

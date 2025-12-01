package main

import (
	"github.com/gin-gonic/gin"

	"kotiki/cats"
	"kotiki/handlers"
	"kotiki/owners"
	"kotiki/users"
)

func main() {
	r := gin.Default()

	// repos

	userRepo := users.NewInMemory()
	catRepo := cats.NewInMemory()
	ownerRepo := owners.NewInMemory()

	// auth

	authService := users.NewAuthService(userRepo)

	// handlers

	userHandlers := handlers.NewUserHandlers(authService)
	catHandlers := handlers.NewCatHandlers(catRepo, ownerRepo)
	ownerHandlers := handlers.NewOwnerHandlers(ownerRepo)

	// users

	r.POST("/register", userHandlers.Register)
	r.POST("/login", userHandlers.Login)
	r.GET("/me", handlers.AuthRequired, userHandlers.Me)

	// owners

	r.GET("/owners", ownerHandlers.GetAllOwners)
	r.GET("/owners/:id", ownerHandlers.GetOwner)
	r.POST("/owners", ownerHandlers.CreateOwner)
	r.DELETE("/owners/:id", ownerHandlers.DeleteOwner)

	// cats

	// список котов + фильтры ?breed=&owner=
	r.GET("/cats", catHandlers.GetAllCat)

	// создать кота (нужна авторизация и внутри CreateCat проверяется, что owner существует)
	r.POST("/cats", handlers.AuthRequired, catHandlers.CreateCat)

	// получить кота по id
	r.GET("/cats/:id", catHandlers.GetCat)

	// удалить кота (только авторизованный пользователь)
	r.DELETE("/cats/:id", handlers.AuthRequired, catHandlers.DeleteCat)

	// чек

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	if err := r.Run(":8080"); err != nil {
		return
	}
}

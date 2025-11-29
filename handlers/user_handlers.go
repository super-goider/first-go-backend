package handlers

import (
	"github.com/gin-gonic/gin"
	"kotiki/users"
)

type UserHandlers struct {
	auth *users.AuthService
}

func NewUserHandlers(auth *users.AuthService) *UserHandlers {
	return &UserHandlers{auth: auth}
}

func (h *UserHandlers) Register(c *gin.Context) {
	var req users.UserCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}
	user, err := h.auth.Register(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)

}

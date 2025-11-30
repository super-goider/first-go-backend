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
func (h *UserHandlers) Login(c *gin.Context) {
	var req users.UserLoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	user, err := h.auth.Login(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sessionID, er := users.CreateSession(user.ID)

	if er != nil {
		c.JSON(400, gin.H{"error": "invalid session"})
		return
	}

	c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)

	c.JSON(200, user)
}

func AuthRequired(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(401, gin.H{"error": "missing session"})
		c.Abort()
		return
	}

	userID, ok := users.GetUserIDBySession(sessionID)
	if !ok {
		c.JSON(401, gin.H{"error": "invalid session"})
		c.Abort()
		return
	}

	c.Set("userID", userID)
	c.Next()
}

func (h *UserHandlers) Me(c *gin.Context) {
	userIDValue, ok := c.Get("userID")
	if !ok {
		c.JSON(500, gin.H{"error": "userID not found in context"})
		return
	}

	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "userID has wrong type"})
		return
	}

	c.JSON(200, gin.H{
		"user_id": userID,
	})
}

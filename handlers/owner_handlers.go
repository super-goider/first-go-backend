package handlers

import (
	"github.com/gin-gonic/gin"
	"kotiki/owners"
	"strconv"
)

type OwnerHandlers struct {
	repo owners.OwnerRepo
}

func NewOwnerHandlers(repo owners.OwnerRepo) *OwnerHandlers {
	return &OwnerHandlers{repo: repo}
}

func (h *OwnerHandlers) CreateOwner(c *gin.Context) {
	var req owners.Owner
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	owner, err := h.repo.Add(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot create owner"})
		return
	}

	c.JSON(201, owner)
}

func (h *OwnerHandlers) GetAllOwners(c *gin.Context) {
	ownersList, err := h.repo.All()
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot load owners"})
		return
	}
	c.JSON(200, ownersList)
}

func (h *OwnerHandlers) GetOwner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	owner, found, err := h.repo.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}
	if !found {
		c.JSON(404, gin.H{"error": "owner not found"})
		return
	}

	c.JSON(200, owner)
}

func (h *OwnerHandlers) DeleteOwner(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	ok, err := h.repo.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}
	if !ok {
		c.JSON(404, gin.H{"error": "owner not found"})
		return
	}

	c.JSON(200, gin.H{"deleted": true})
}

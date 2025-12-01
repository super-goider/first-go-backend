package handlers

import (
	"github.com/gin-gonic/gin"
	"kotiki/cats"
	"kotiki/owners"
	"strconv"
)

type CatHandlers struct {
	catRepo   cats.CatRepo
	ownerRepo owners.OwnerRepo
}

func NewCatHandlers(catRepo cats.CatRepo, ownerRepo owners.OwnerRepo) *CatHandlers {
	return &CatHandlers{
		catRepo:   catRepo,
		ownerRepo: ownerRepo,
	}
}

func (h *CatHandlers) CreateCat(c *gin.Context) {
	var req cats.CatCreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	// 1. Проверяем, что владелец существует
	_, found, err := h.ownerRepo.Get(req.OwnerId)
	if err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}
	if !found {
		c.JSON(400, gin.H{"error": "owner not found"})
		return
	}

	// 2. Если владелец есть — создаём кота
	cat := cats.Cat{
		Name:    req.Name,
		Breed:   req.Breed,
		Age:     req.Age,
		OwnerId: req.OwnerId,
		About:   req.About,
	}

	saved, err := h.catRepo.Add(cat)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot save cat"})
		return
	}

	c.JSON(201, saved)
}

func (h *CatHandlers) GetAllCat(c *gin.Context) {

	breed := c.Query("breed")
	owner := c.Query("owner")

	if breed == "" && owner == "" {

		catList, err := h.catRepo.All()

		if err != nil {
			c.JSON(500, gin.H{"error": "invalid cat list"})
			return
		}
		c.JSON(200, catList)
		return
	} else {

		catList, err := h.catRepo.Filter(breed, owner)
		if err != nil {
			c.JSON(500, gin.H{"error": "invalid cat list"})
			return
		}

		c.JSON(200, catList)
		return
	}

}

func (h *CatHandlers) GetCat(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "bad id"})
		return
	}

	resCat, isGot, err := h.catRepo.Get(idInt)

	if err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	if isGot {
		c.JSON(200, resCat)
		return
	} else {
		c.JSON(404, gin.H{"error": "cat does not exist"})
		return
	}

}

func (h *CatHandlers) DeleteCat(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "bad id"})
		return
	}

	isDeleted, err := h.catRepo.Delete(idInt)

	if err != nil {
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	if isDeleted {
		c.Status(204)
		return
	} else {
		c.JSON(404, gin.H{"error": "cat is not deleted"})
		return
	}

}

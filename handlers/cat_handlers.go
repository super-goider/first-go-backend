package handlers

import (
	"github.com/gin-gonic/gin"
	"kotiki/cats"
	"strconv"
)

type CatHandlers struct {
	repo cats.CatRepo // без * потому что интерфейс это уже ссылка
}

func NewCatHandlers(repo cats.CatRepo) *CatHandlers {
	return &CatHandlers{
		repo: repo,
	}
}

func (h *CatHandlers) CreateCat(c *gin.Context) {
	var cat cats.Cat

	if err := c.BindJSON(&cat); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	savedCat, err := h.repo.Add(cat)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot save cat"})
		return
	}
	c.JSON(201, savedCat)
}

func (h *CatHandlers) GetAllCat(c *gin.Context) {

	breed := c.Query("breed")
	owner := c.Query("owner")

	if breed == "" && owner == "" {

		catList, err := h.repo.All()

		if err != nil {
			c.JSON(500, gin.H{"error": "invalid cat list"})
			return
		}
		c.JSON(200, catList)
		return
	} else {

		catList, err := h.repo.Filter(breed, owner)
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

	resCat, isGot, err := h.repo.Get(idInt)

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

	isDeleted, err := h.repo.Delete(idInt)

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

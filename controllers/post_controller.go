package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-like-twitter-api/form/api"
	"github.com/katsuomi/gin-like-twitter-api/models"
	"github.com/katsuomi/gin-like-twitter-api/models/repository"
)

// Controller is user controlller
type PostController struct{}

// Index action: GET /posts
func (pc PostController) Index(c *gin.Context) {
	var u repository.PostRepository
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /posts
func (pc PostController) Create(c *gin.Context) {
	var u repository.PostRepository
	in := api.Post{}
	if err := c.BindJSON(&in); err != nil {
		return
	}
	in2 := &models.Post{
		ID:      in.ID,
		Content: in.Content,
		UserID:  in.UserID,
	}
	p, err := u.CreateModel(in2)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}

// Show action: Get /posts/:id
func (pc PostController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var p repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	post, err := p.GetByID(idInt)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, post)
	}
}

// Update action: Put /posts/:id
func (pc PostController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	p, err := u.UpdateByID(idInt, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /posts/:id
func (pc PostController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	if err := u.DeleteByID(idInt); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "ID" + id + "の投稿を削除しました"})
	return
}

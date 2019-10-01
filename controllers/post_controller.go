package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-gorm-twitter-api/form/api"
	"github.com/katsuomi/gin-gorm-twitter-api/models"
	"github.com/katsuomi/gin-gorm-twitter-api/models/repository"
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

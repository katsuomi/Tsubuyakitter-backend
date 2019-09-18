package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-gorm-practice/service"
)

// Controller is post controlller
type PostController struct{}

// Index action: GET /posts
func (pc PostController) Index(c *gin.Context) {
	var s service.PostService
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /posts
func (pc PostController) Create(c *gin.Context) {
	var s service.PostService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /posts/:id
func (pc PostController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.PostService
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /posts/:id
func (pc PostController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.PostService
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /posts/:id
func (pc PostController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.PostService

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

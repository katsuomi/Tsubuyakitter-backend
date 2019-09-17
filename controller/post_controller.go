package post

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"../service"
)

// Controller is post controlller
type Controller struct{}

// Index action: GET /posts
func (pc Controller) Index(c *gin.Context) {
	var s post.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /posts
func (pc Controller) Create(c *gin.Context) {
	var s post.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /posts/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s post.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /posts/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s post.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /posts/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s post.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
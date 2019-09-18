package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-gorm-practice/service"
)

// Controller is user controlller
type UserController struct{}

// Index action: GET /users
func (pc UserController) Index(c *gin.Context) {
	var s service.UserService
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (pc UserController) Create(c *gin.Context) {
	var s service.UserService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (pc UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (pc UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

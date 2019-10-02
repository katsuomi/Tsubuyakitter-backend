package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-gorm-twitter-api/db"
	"github.com/katsuomi/gin-gorm-twitter-api/form/api"
	"github.com/katsuomi/gin-gorm-twitter-api/models"
)

// Service procides post's behavior
type PostRepository struct{}

// Post is alias of entity.Post struct
type Post api.Post

// GetAll is get all Post
func (_ PostRepository) GetAll() ([]Post, error) {
	db := db.GetDB()
	var p []Post

	if err := db.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

// CreateModel is create Post model
func (_ PostRepository) CreateModel(p *models.Post) (*models.Post, error) {
	fmt.Printf("%+v", p)
	db := db.GetDB()
	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// GetByID is get a Post by ID
func (_ PostRepository) GetByID(id int) (models.Post, error) {
	db := db.GetDB()
	var p models.Post
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// UpdateByID is update a Post
func (_ PostRepository) UpdateByID(id int, c *gin.Context) (api.Post, error) {
	db := db.GetDB()
	var p api.Post
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	userID := p.UserID
	if err := c.BindJSON(&p); err != nil {
		return p, err
	}
	fmt.Printf("%+V", p)
	p.ID = uint(id)
	p.UserID = userID
	db.Save(&p)

	return p, nil
}

// DeleteByID is delete a Post by ID
func (_ PostRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var p Post

	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}

	return nil
}

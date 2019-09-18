package service

import (
	"github.com/gin-gonic/gin"

	"github.com/katsuomi/gin-gorm-practice/db"
	"github.com/katsuomi/gin-gorm-practice/entity"
)

// Service procides post's behavior
type PostService struct{}

// Post is alias of entity.Post struct
type Post entity.Post

// GetAll is get all Post
func (s PostService) GetAll() ([]Post, error) {
	db := db.GetDB()
	var u []Post

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// CreateModel is create Post model
func (s PostService) CreateModel(c *gin.Context) (Post, error) {
	db := db.GetDB()
	var u Post

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a Post
func (s PostService) GetByID(id string) (Post, error) {
	db := db.GetDB()
	var u Post

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateByID is update a Post
func (s PostService) UpdateByID(id string, c *gin.Context) (Post, error) {
	db := db.GetDB()
	var u Post

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a Post
func (s PostService) DeleteByID(id string) error {
	db := db.GetDB()
	var u Post

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}

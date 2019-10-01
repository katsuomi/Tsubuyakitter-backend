package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/katsuomi/gin-gorm-twitter-api/db"
	"github.com/katsuomi/gin-gorm-twitter-api/models"
)

// Service procides post's behavior
type PostRepository struct{}

// Post is alias of entity.Post struct
type Post models.Post

type CreatePost struct {
	ID      uint
	Content string
	UserID  uint
	User    *gorm.DB
}

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

package models

// Post is post models property
type Post struct {
	ID      uint   `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
	User    User   `json:"-" binding:"required"`
	UserID  uint   `gorm:"not null"  json:"user_id"`
}

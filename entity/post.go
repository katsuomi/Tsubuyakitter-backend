package entity

// Post is post models property
type Post struct {
	gorm.Model
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	UserRefer uint
}
//gitgit 
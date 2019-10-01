package models

// User is user models property
type User struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Posts []Post `json:"posts"`
}

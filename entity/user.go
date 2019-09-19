package entity

// User is user models property
type User struct {
	gorm.Model
	Posts []Post
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

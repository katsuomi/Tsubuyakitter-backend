package api

// swagger:model User
type User struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// swagger:model User
type UserPosts struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Posts []Post `json:"posts"`
}

package entity

// Post is post models property
type Post struct {
	ID        uint   `json:"id"`
	USER_ID   uint   `json:"user_id"`
	content   string `json:"content"`
}

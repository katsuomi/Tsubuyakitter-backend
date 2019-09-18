package entity

// Post is post models property
type Post struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	UserId    uint   `json:"user_id"`
}
//git
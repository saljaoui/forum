package posts

import (
	"forum-project/backend/internal/repository/categories"
	"time"
)

type Post struct {
	Id         int64     `json:"id"`
	User_id    int64     `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Category   string    `json:"category"`
	Created_at time.Time `json:"created_at"`
}

func (p *Post) AddPost() {
	insertPost(p)
	categories.InsertCategory(p.Id, p.Category)
}

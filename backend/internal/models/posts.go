package models

import "time"

type Post struct {
	Id        int64     `json:"id"`
	User_id int64    `json:"user_id"`
	Title  string    `json:"title"`
	Content     string    `json:"content"`
	CreatedAt time.Time `json:"createdat"`
}

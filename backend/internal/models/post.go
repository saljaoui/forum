package models

import "time"

type Post struct {
	Id         int64     `json:"id"`
	User_id    int64     `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Category   string    `json:"category"`
	Created_at time.Time `json:"created_at"`
}
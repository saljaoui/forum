package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdet"`
}

type Post struct {
	Id       int
	User_id  int
	Title    string
	Content  string
	Category string
}

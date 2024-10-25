package models

import "time"

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Post struct {
	ID      int64
	User_id int64
	Title   string
	Content string
	Category string
	Created_at time.Time
}

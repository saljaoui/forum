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

	Id      int
	User_id int
	Title   string
	Content string
}

package home

import "time"

type PostResponde struct {
	Card_Id           int
	Post_Id      int
	UserID       int
	FirstName    string
	LastName     string
	Title        string
	Content      string
	CategoryName string
	Likes        int
	CreatedAt    time.Time
}


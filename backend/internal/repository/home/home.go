package home

import "time"

type PostResponde struct {
	ID           int
	UserID       int
	FirstName    string
	LastName     string
	Title        string
	Content      string
	CategoryName string
	CreatedAt    time.Time
}

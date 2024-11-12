package models

import "time"

type User struct {
	Id        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UUID      time.Time `json:"uuid"`
}

type Login struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

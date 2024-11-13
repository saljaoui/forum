package models

type card struct {
	Id  int
	User_Id  int
	Content   string
	CreatedAt string
}

func (c *card) NewCard(user_id int,content string) *card {
	return &card{
		Id: -1,
		User_Id: user_id,
		Content: content,
	}
}

func (c *card) Add() int {
	return 0
}
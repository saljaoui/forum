package comment

import "forum-project/backend/internal/repository/cards"

type Comment struct {
	ID        int    `json:"id"`
	User_Id   int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	Card_Id   int    `json:"card_id"`
	Target_Id int    `json:"target_id"`
}

func NewComment(user_id int, content string, target int) *Comment {
	return &Comment{
		ID:        -1,
		Card_Id:   -1,
		Target_Id: target,
		User_Id:   user_id,
		Content:   content,
	}
}

func (c *Comment) Add() int {
	card := cards.NewCard(c.User_Id, c.Content)
	card.Add()
	if card.Id == -1 {
		return -1
	}
	c.Card_Id = card.Id
	c.ID = insertComment(c.Card_Id, c.Target_Id)
	return c.ID
}



func GetComment(id int) *Comment {
	data_Row := getCommentById(id)
	if data_Row == nil{
		return nil
	}
	card := cards.GetCard(data_Row.Card_Id)
	if card == nil {
		return nil
	}
	newComment := &Comment{
		ID: data_Row.ID,
		Card_Id: data_Row.Card_Id,
		Target_Id: data_Row.Target_Id,
		Content: card.Content,
		User_Id: card.User_Id,
		CreatedAt: card.CreatedAt,
	}
	return newComment
}









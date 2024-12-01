package comment

import (
	"forum-project/backend/internal/repository/cards"
)

type Comment struct {
	ID        int    `json:"id"`
	User_Id   int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	Card_Id   int    `json:"card_id"`
	Target_Id int    `json:"target_id"`
}

type comment_View struct {
	Comment_ID int 		`json:"commentid"`
	User_Id    int 		`json:"userid"`
	Firstname  string	`json:"first_name"`
	Lastname   string	`json:"last_name"`
	Content    string	`json:"content"`
	CreatedAt  string	`json:"createdat"`
	Card_Id    int		`json:"cardid"`
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
	if data_Row == nil {
		return nil
	}
	card := cards.GetCard(data_Row.Card_Id)
	if card == nil {
		return nil
	}
	newComment := &Comment{
		ID:        data_Row.ID,
		Card_Id:   data_Row.Card_Id,
		Target_Id: data_Row.Target_Id,
		Content:   card.Content,
		User_Id:   card.User_Id,
		CreatedAt: card.CreatedAt,
	}
	return newComment
}

func GetAllCommentsbyTarget(target int) []comment_View {
	list_Comments := make([]comment_View, 0)
	list := getAllCommentsbyTargetId(target)
	size := len(list)
	if size == 0 {
		return nil
	}
	for index := 0; index < size; index++ {
		comment := convert(list[index])
		list_Comments = append(list_Comments, comment)
	}
	return list_Comments
}

func convert(cmrv comment_Row_View) comment_View {
	comment := comment_View{}
	comment.Comment_ID = cmrv.comment_ID
	comment.User_Id = cmrv.User_Id
	comment.Firstname = cmrv.firstname
	comment.Lastname = cmrv.lastname
	comment.Content = cmrv.Content
	comment.CreatedAt = cmrv.CreatedAt
	comment.Card_Id = cmrv.Card_Id
	return comment
}

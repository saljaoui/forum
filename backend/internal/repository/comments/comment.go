package comment

type comment struct {
	ID        int    `json:"id"`
	User_Id   int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	Card_Id   int    `json:"card_id"`
	Target_Id int    `json:"target_id"`
}

func NewComment(user_id int, content string, target int) *comment {
	return &comment{
		ID:        -1,
		Card_Id:   -1,
		Target_Id: target,
		User_Id:   user_id,
		Content:   content,
	}
}

// func (c *comment) Add() int {
// 	card := cards.NewCard(c.User_Id, c.Content)
// 	card.Add()
// 	if card.Id == -1 {
// 		return -1
// 	}
// 	c.Card_Id = card.Id
// 	return insertComment(c.Card_Id, c.Target_Id)
// }

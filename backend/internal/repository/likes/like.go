package like

import "errors"

type like struct {
	ID      int `json:"id"`
	User_Id int `json:"user_id"`
	Card_Id int `json:"card_id"`
	isLiked int `json:"is_liked"`
}

func NewLike(user_id, card_id int) *like {
	return &like{
		ID:      -1,
		User_Id: user_id,
		Card_Id: card_id,
		isLiked: -1,
	}
}
func (l *like) SetIsLike(val int) error {
	if val < -1 || val > 1 {
		return errors.New("like is a property that can be -1, 0 or 1 .")
	}
	l.isLiked = val
	return nil
}

func (l *like) GetIsLike() int {
	return l.isLiked
}
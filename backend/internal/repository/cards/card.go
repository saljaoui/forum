package cards

import "fmt"

type card struct {
	Id  int
	User_Id  int
	Content   string
	CreatedAt string
}

func NewCard(user_id int , content string) *card {
	return &card{
		Id: -1,
		User_Id: user_id,
		Content: content,
	}
}

func (C *card) Add() int {
	return insertCard(C.User_Id,C.Content)
}

func GetCard(id int) *card {
	myCard := card{
		Id: id,
	}
	card_Row := getCardById(id)
	if card_Row != nil {
		myCard.User_Id = card_Row.User_Id
		myCard.Content = card_Row.Content
		myCard.CreatedAt = card_Row.CreatedAt
		return &myCard
	}
	return nil
}

func (C *card) PrintInfo() {
	if C == nil {
		fmt.Println("this card is NULL")
		return 
	}
	fmt.Printf("id : %d \n",C.Id)
	fmt.Printf("user_ID : %d \n",C.User_Id)
	fmt.Printf("content : %s \n",C.Content)
	fmt.Printf("date : %s \n",C.CreatedAt)
}
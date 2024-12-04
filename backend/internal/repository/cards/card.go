package cards

import "fmt"

type card struct {
	Id  int
	User_Id  int
	Content   string
	CreatedAt string
}


type card_View struct {
	Id  int				`json:"id"`
	User_Id  int		`json:"userid"`
	Content   string	`json:"content"`
	CreatedAt string	`json:"date"`
	FirstName string	`json:"firstName"`
	LastName  string	`json:"lastName"`
	Likes 	  int		`json:"likes"`
	DisLikes  int		`json:"dislikes"`
	Comments  int		`json:"comments"`

}

func NewCard(user_id int , content string) *card {
	return &card{
		Id: -1,
		User_Id: user_id,
		Content: content,
	}
}

func (C *card) Add() int {
	C.Id = insertCard(C.User_Id,C.Content)
	return C.Id
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

func GetOneCard(id int) card_View {
	card := card_View{}
	data_Row := getCard(id)
	card.Id = data_Row.Id
	card.User_Id = data_Row.User_Id
	card.Content = data_Row.Content
	card.CreatedAt = data_Row.CreatedAt
	card.FirstName = data_Row.FirstName
	card.LastName = data_Row.LastName
	card.Likes = data_Row.Likes
	card.DisLikes = data_Row.DisLikes
	card.Comments = data_Row.Comments
	return card
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
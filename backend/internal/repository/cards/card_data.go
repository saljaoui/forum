package cards

import (
	//"fmt"
	"forum-project/backend/internal/database"
)

type Card_Row struct {
	Id  int
	User_Id  int
	Content   string
	CreatedAt string
}

func insertCard(user_id int , content string) int {
    query := "INSERT INTO card(user_id,content) VALUES(?,?)"
    return database.Exec(query,user_id,content)    
}

func getCardById(id int) *Card_Row {
    query := "SELECT * FROM card WHERE card.id =?;"
    myCard_Row := &Card_Row{}
    err := database.SelectOneRow(query,id).Scan(&id,&myCard_Row.User_Id,&myCard_Row.Content,&myCard_Row.CreatedAt)

    if err != nil {
        return nil
    } else {
        return myCard_Row
    }
}
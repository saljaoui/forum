package comment

import "forum-project/backend/internal/database"


type comment_Row struct {
	ID        int
	User_Id   int
	Content   string
	CreatedAt string
	Card_Id   int
	Target_Id int
}

func insertComment(card_id,target_id int) int {
    query := "INSERT INTO comment(card_id,target_id) VALUES(?,?);"
    return database.Exec(query,card_id,target_id)
}

func GetCommentById(id int) *comment_Row {
    Row := comment_Row{}
    query := "SELECT * FROM comment WHERE comment.id =?;"
    err := database.SelectOneRow(query,Row.ID,Row.Card_Id,Row.Target_Id)
    if err != nil{
        return nil
    }
    return &Row
}
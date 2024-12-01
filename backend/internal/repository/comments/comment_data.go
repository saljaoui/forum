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

type comment_Row_View struct {
	comment_ID int
	User_Id    int
	firstname  string
	lastname   string
	Content    string
	CreatedAt  string
	Card_Id    int
}

func insertComment(card_id, target_id int) int {
	query := "INSERT INTO comment(card_id,target_id) VALUES(?,?);"
	resl, _ := database.Exec(query, card_id, target_id)
	id, err := resl.LastInsertId()
	if err != nil {
		return -1
	}
	return int(id)
}

func getCommentById(id int) *comment_Row {
    Row := comment_Row{}
    query := "SELECT * FROM comment WHERE comment.id =?;"
    err := database.SelectOneRow(query,id).Scan(&Row.ID,&Row.Card_Id,&Row.Target_Id)
    if err != nil{
        return nil
    }
    return &Row 
}

func getAllCommentsbyTargetId(target int) []comment_Row_View {
	list_Comments := make([]comment_Row_View, 0)
	query := `SELECT cm.id,u.id,u.firstname,u.lastname,c.content,c.created_at,c.id from comment cm JOIN card c
              ON c.id = cm.card_id JOIN user u on c.user_id = u.id 
              WHERE cm.target_id =?;`
	data_Rows := database.SelectRows(query,target)
    for data_Rows.Next(){
        Row := comment_Row_View{}
        err := data_Rows.Scan(&Row.comment_ID,&Row.User_Id,&Row.firstname,&Row.lastname,&Row.Content,&Row.CreatedAt,&Row.Card_Id)
        if err != nil {
            return nil
        }    
        list_Comments = append(list_Comments, Row)
    }
	return list_Comments
}

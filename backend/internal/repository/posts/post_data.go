package posts

import (
	"fmt"

	"forum-project/backend/internal/database"
)

<<<<<<< HEAD
func inserPost(title string, card_id int) int64 {
	query := "INSERT INTO post(title, card_id) VALUES(?,?);"
	row, err := database.Exec(query, title, card_id)
	if err != nil {
		fmt.Println("error to insert")
	}
	id, err := row.LastInsertId()
	if err != nil {
		fmt.Println("Error to get id ")
	}
	return id
}
=======
func AddPost(post *Post) messages.Messages {
    message := messages.Messages{}
    query := "INSERT INTO card (user_id,content) VALUES(?,?)"
    database.Exec(query, post.User_id, post.Content)
    categorie.PostCategory(post.Id, post.Category)
    return message
}
>>>>>>> 0cd3f2d5e1988ac649f4af840c0c9655f88055e9

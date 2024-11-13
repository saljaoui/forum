package repository

import (
	"fmt"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"
)

func Post(post *models.Post) messages.Messages {
	message := messages.Messages{}
	fmt.Println(post.User_id, post.Content)
	insert := "INSERT INTO card (user_id,content) VALUES(?,?)"

	database.Exec(insert, post.User_id, post.Content)

	return message
}

package posts

import (
	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
)

func insertPost(post *Post) messages.Messages {
	message := messages.Messages{}
	query := "INSERT INTO card (user_id,content) VALUES(?,?)"
	database.Exec(query, post.User_id, post.Content)
	return message
}

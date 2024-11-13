package posts

import (
	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"
	"forum-project/backend/internal/repository/categories"
)

func Post(post *models.Post) messages.Messages {
	message := messages.Messages{}
	query := "INSERT INTO card (user_id,content) VALUES(?,?)"
	category.PostCategory(post)
	database.Exec(query, post.User_id, post.Content)
	return message
}

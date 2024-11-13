<<<<<<<< HEAD:backend/internal/models/posts/post_data.go
package models
========
package posts
>>>>>>>> soufian:backend/internal/repository/posts/post_data.go

import (
	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
<<<<<<<< HEAD:backend/internal/models/posts/post_data.go
	models "forum-project/backend/internal/models/categories"
========
	"forum-project/backend/internal/models"
	"forum-project/backend/internal/repository/categories"
>>>>>>>> soufian:backend/internal/repository/posts/post_data.go
)

func Post(post *models.Post) messages.Messages {
	message := messages.Messages{}
	query := "INSERT INTO card (user_id,content) VALUES(?,?)"
	category.PostCategory(post)
	database.Exec(query, post.User_id, post.Content)
	return message
}

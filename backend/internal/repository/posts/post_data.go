package posts

import (
    messages "forum-project/backend/internal/Messages"
    "forum-project/backend/internal/database"
    "forum-project/backend/internal/repository/categories"
)

func AddPost(post *Post) messages.Messages {
    message := messages.Messages{}
    query := "INSERT INTO card (user_id,content) VALUES(?,?)"
    database.Exec(query, post.User_id, post.Content)
    categories.PostCategory(post.Id, post.Category)
    return message
}

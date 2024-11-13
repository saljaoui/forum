package repository

import (
	"fmt"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"
)

func Post(post *models.Post) messages.Messages {
	message := messages.Messages{}
	query := "INSERT INTO card (user_id,content) VALUES(?,?)"
	PostCategory(post)
	database.Exec(query, post.User_id, post.Content)
	return message
}

func PostCategory(post *models.Post) {
	categoryId := GetCategoryId(post)
	fmt.Println("oooooooooooooo",categoryId, post.Id)
	query := "INSERT INTO post_category (post_id, category_id) VALUES(?,?)"
	database.Exec(query, post.Id, categoryId)
}

func GetCategoryId(post *models.Post) int {
	query := "SELECT id FROM category WHERE name = ?"
	db := database.Config()
	categoryId := 0
	db.QueryRow(query, post.Category).Scan(&categoryId)
    return categoryId
}

package category

import (
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"
)

func PostCategory(post *models.Post) {
	categoryId := GetCategoryId(post)
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

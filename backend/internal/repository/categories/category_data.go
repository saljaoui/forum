package categories

import (
	"forum-project/backend/internal/database"
)

func postCategory(postId int64, category string) {
	categoryId := getCategoryId(category)
	query := "INSERT INTO post_category (post_id, category_id) VALUES(?,?)"
	database.Exec(query, postId, categoryId)
}

func getCategoryId(category string) int {
	categoryId := 0
	query := "SELECT id FROM category WHERE name = ?"
	database.SelectOneRow(query, category).Scan(&categoryId)
	return categoryId
}

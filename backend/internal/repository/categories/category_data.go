package categories

import (
    "forum-project/backend/internal/database"
)

func PostCategory(postId int64, category string) {
    categoryId := GetCategoryId(category)
    query := "INSERT INTO post_category (post_id, category_id) VALUES(?,?)"
    database.Exec(query, postId, categoryId)
}

func GetCategoryId(category string) int {
    categoryId := 0
    query := "SELECT id FROM category WHERE name = ?"
    db := database.Config()
    db.QueryRow(query, category).Scan(&categoryId)
    return categoryId
}
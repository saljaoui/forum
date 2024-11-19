package category

import (
	"fmt"

	"forum-project/backend/internal/database"
)

func postCategory(postId int, category string) {
	categoryId := getCategoryId(category)
	query := "INSERT INTO post_category (post_id, category_id) VALUES(?,?)"
	_, err := database.Exec(query, postId, categoryId)
	if err != nil {
		fmt.Println(err)
	}
}

func getCategoryId(category string) int {
	categoryId := -1
	query := "SELECT id FROM category WHERE name = ?"
	err:=database.SelectOneRow(query, category).Scan(&categoryId)
	if err!=nil {
		fmt.Println("error To Get Id") 
	}
	return categoryId
}

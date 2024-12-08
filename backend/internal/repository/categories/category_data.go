package category

import (
	"fmt"
	"forum-project/backend/internal/database"
)

func postCategory(postId int, category string) error {
	categoryId, err := getCategoryId(category)
	// fmt.Println(categoryId)
	// fmt.Println(categoryId,postId)
	if err != nil {
		return err
	}
	query := "INSERT INTO post_category (post_id, category_id) VALUES(?,?)"
	_, err = database.Exec(query, postId, categoryId)
	if err != nil {

		return err
	}
	return nil
}

func getCategoryId(category string) (int, error) {
	categoryId := 0
	query := "SELECT id FROM category WHERE name = ?"
	db := database.Config()
	err := db.QueryRow(query, category).Scan(&categoryId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return categoryId, nil
}

// func getPostsByCategoryId(category int) {
// 	query := "SELECT * FROM post_category pc, post p WHERE pc.post_id = p.id AND pc.category_id = 3"

// }
package posts

import (
	"fmt"

	"forum-project/backend/internal/database"
	like "forum-project/backend/internal/repository/likes"
)

func inserPost(title string, card_id int) int64 {
	query := "INSERT INTO post(title, card_id) VALUES(?,?);"
	row, err := database.Exec(query, title, card_id)
	if err != nil {
		fmt.Println("error to insert")
	}
	id, err := row.LastInsertId()
	if err != nil {
		fmt.Println("Error to get id ")
	}
	return id
}

func GetPosts(query string) []PostResponde {
	db := database.Config()
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var posts []PostResponde
	for rows.Next() {
		var post PostResponde
		err := rows.Scan(
			&post.Card_Id,
			&post.Post_Id,
			&post.UserID,
			&post.FirstName,
			&post.LastName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			return nil
		}
		likes, dislikes := like.GetLikes(post.Post_Id)
		post.Likes = likes
		post.Dislikes = dislikes
		posts = append(posts, post)
	}
	return posts
}

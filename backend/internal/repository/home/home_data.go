package home

import (
	"fmt"
	"strconv"

	"forum-project/backend/internal/database"
)

func GetPosts(quantity int) []PostResponde {
	query := `SELECT p.card_id AS 'card_id', p.id, u.id AS 'user_id', u.firstname, u.lastname, p.title, c.content, cat.name, c.created_at  
	FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id 
	AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id`
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
			&post.CategoryName,
			&post.CreatedAt,
		)
		if err != nil {
			return nil
		}
		likes := getLikes(post.Post_Id)
		post.Likes = likes
		fmt.Println(likes)
		posts = append(posts, post)
	}
	return posts
}

func getLikes(post_id int) int {
	query := `SELECT sum(is_like) FROM post p, likes l WHERE p.card_id = l.card_id AND p.id = ` + strconv.Itoa(post_id)
	likes := 0
	err := database.SelectOneRow(query).Scan(&likes)
	if err != nil {
		return 0
	}
	return likes
}

// func canPurchase(id int, quantity int) (bool, error) {
//     var enough bool
//     // Query for a value based on a single row.
//     if err := db.QueryRow("SELECT (quantity >= ?) from album where id = ?",
//         quantity, id).Scan(&enough); err != nil {
//         if err == sql.ErrNoRows {
//             return false, fmt.Errorf("canPurchase %d: unknown album", id)
//         }
//         return false, fmt.Errorf("canPurchase %d: %v", id, err)
//     }
//     return enough, nil
// }

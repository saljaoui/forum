package home

import (
	"forum-project/backend/internal/database"
)

func GetPosts(quantity int) []PostResponde {
	query := `SELECT p.id, u.id AS 'user_id', u.firstname, u.lastname, p.title, c.content, cat.name, c.created_at  
	FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id 
	AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id AND cat.name="Technology"`
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
			&post.ID,
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
		posts = append(posts, post)
	}
	return posts
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

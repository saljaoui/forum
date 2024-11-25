package home

import (
	"strings"

	"forum-project/backend/internal/database"
)

func GetPosts(quantity int) []PostResponde {
	query := `
SELECT
    p.id,
    u.id AS 'UserID',
    u.firstname,
    u.lastname,
    p.title,
    c.content,
    GROUP_CONCAT(cat.name),
    c.created_at
FROM
    post p
    JOIN card c ON p.card_id = c.id
    JOIN user u ON c.user_id = u.id
    JOIN post_category pc ON p.id = pc.post_id
    JOIN category cat ON pc.category_id = cat.id
GROUP BY
    p.id,
    u.id,
    u.firstname,
    u.lastname,
    p.title,
    c.content,
    c.created_at
	ORDER BY
    c.created_at DESC;`
	db := database.Config()
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var posts []PostResponde
	for rows.Next() {
		var categoryNames string
		var post PostResponde
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.FirstName,
			&post.LastName,
			&post.Title,
			&post.Content,
			&categoryNames,
			&post.CreatedAt,
		)
		if err != nil {
			return nil
		}
		post.CategoryName = strings.Split(categoryNames, ",")
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

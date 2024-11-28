package home

import (
	"strconv"

	"forum-project/backend/internal/database"
)

func GetPosts(categoryId int) []PostResponde {
	query := `SELECT p.card_id AS 'card_id', p.id, u.id AS 'user_id', u.firstname, u.lastname, p.title, c.content, c.created_at  
	FROM post p, card c, user u WHERE p.card_id=c.id 
	AND c.user_id=u.id`
	if categoryId != 0 {
		query = `SELECT p.card_id AS 'card_id', p.id, u.id AS 'user_id', u.firstname, u.lastname, p.title, c.content, c.created_at  
	FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id 
	AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id AND cat.id =` + strconv.Itoa(categoryId)
	}
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
		likes, dislikes := getLikes(post.Post_Id)
		post.Likes = likes
		post.Dislikes = dislikes
		posts = append(posts, post)
	}
	return posts
}

func getLikes(post_id int) (int, int) {
	querylike := `SELECT sum(is_like) FROM post p, likes l WHERE p.card_id = l.card_id AND l.is_like = 1 AND p.id = ` + strconv.Itoa(post_id)
	like := 0
	err := database.SelectOneRow(querylike).Scan(&like)
	if err != nil {
		like = 0
	}
	querydislike := `SELECT sum(is_like) FROM post p, likes l WHERE p.card_id = l.card_id AND l.is_like = -1 AND p.id = ` + strconv.Itoa(post_id)
	dislike := 0
	err = database.SelectOneRow(querydislike).Scan(&dislike)
	if err != nil {
		dislike = 0
	}
	return like, dislike * -1
}

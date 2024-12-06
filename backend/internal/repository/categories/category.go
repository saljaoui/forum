package category

import (
	post "forum-project/backend/internal/repository/posts"
)

type Category struct {
	Category string `json:"Category"`
}

func AddCategory(post_id int, category string) error {
	err := postCategory(post_id, category)
	if err != nil {
		return err
	}
	return nil
}

func GetPostsByCategoryId(categoryName string) []post.PostResponde {
	query := `SELECT p.card_id AS 'card_id', p.id, u.id AS 'user_id', u.firstname, u.lastname, p.title, c.content, c.created_at
	FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id
	AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id AND cat.name = "` + categoryName + "\""
	return post.GetPosts(query)
}

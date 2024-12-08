package profile

import (
	"forum-project/backend/internal/repository/posts"
	"strconv"
)

func GetPostsProfile(user_id int) []posts.PostResponde {
	query := `SELECT p.card_id AS 'card_id', p.id, u.id AS 'user_id', u.firstname, u.lastname, p.title, c.content, c.created_at  
	FROM post p, card c, user u WHERE p.card_id=c.id 
	AND c.user_id=u.id AND u.id = ` +strconv.Itoa(user_id)
	return posts.GetPosts(query)
}

func GetPostsProfileByLikes(user_id int) []posts.PostResponde {
	query := `SELECT p.card_id AS 'card_id', p.id, u.id AS 'user_id', u.firstname, u.lastname, p.title, c.content, c.created_at  
	FROM post p, card c, user u, likes l WHERE p.card_id=c.id 
	AND c.user_id=u.id AND p.card_id = l.card_id AND l.user_id = ` +strconv.Itoa(user_id)
	return posts.GetPosts(query)
}


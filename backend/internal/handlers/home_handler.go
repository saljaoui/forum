package handlers

import (
	"encoding/json"
	"forum-project/backend/internal/repository/home"
	"net/http"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	posts := home.GetPosts(5)
	json.NewEncoder(w).Encode(posts)
	//SELECT p.id, u.id, u.firstname, u.lastname, p.title, c.content, cat.name, c.created_at  FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id
}

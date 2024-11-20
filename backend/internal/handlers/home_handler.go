package handlers

import (
	"net/http"

	"forum-project/backend/internal/repository/home"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	posts := home.GetPosts(2)
	newEncode := NewEncoderJsone(w)
	err := newEncode.Encode(&posts)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
	}
	// SELECT p.id, u.id, u.firstname, u.lastname, p.title, c.content, cat.name, c.created_at
	// FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id AND
	// c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id
}

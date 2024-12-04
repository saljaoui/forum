package handlers

import (
	"encoding/json"
	"net/http"

	"forum-project/backend/internal/repository/cards"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, "Method Not Allowd", http.StatusMethodNotAllowed)
	}
	posts := cards.GetAllCards()
	json.NewEncoder(w).Encode(posts)
	// SELECT p.id, u.id, u.firstname, u.lastname, p.title, c.content, cat.name, c.created_at  FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id
}

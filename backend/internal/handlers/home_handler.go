package handlers

import (
	"encoding/json"
	"net/http"

	"forum-project/backend/internal/repository/cards"
	like "forum-project/backend/internal/repository/likes"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}
	posts := cards.GetAllCards()
	json.NewEncoder(w).Encode(posts)
	// SELECT p.id, u.id, u.firstname, u.lastname, p.title, c.content, cat.name, c.created_at  FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id
}

func LikesHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsoneResponse(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}
	liked := like.Like{}
	decode := DecodeJson(r)
	err := decode.Decode(&liked)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	dislike := liked.ChecklikesUser()
	// fmt.Println(dislike)
	json.NewEncoder(w).Encode(dislike)
	// JsoneResponse(w, dislike, http.StatusOK)
}

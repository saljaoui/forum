package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum-project/backend/internal/repository/home"
	like "forum-project/backend/internal/repository/likes"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, "Method Not Allowd", http.StatusMethodNotAllowed)
	}
	posts := home.GetPostsHome()
	json.NewEncoder(w).Encode(posts)
	// SELECT p.id, u.id, u.firstname, u.lastname, p.title, c.content, cat.name, c.created_at  FROM post p, card c, user u, post_category pc, category cat WHERE p.card_id=c.id AND c.user_id=u.id AND p.id = pc.post_id AND pc.category_id=cat.id
}

func LikesHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsoneResponse(w, "Method Not Allowd", http.StatusMethodNotAllowed)
	}
	liked := like.Like{}
	decode := DecodeJson(r)
	err := decode.Decode(&liked)
	if err != nil {
		JsoneResponse(w, err.Error(), http.StatusMethodNotAllowed)
	}
	lik, dislike := liked.ChecklikesUser(1, -1)
	fmt.Println(lik, dislike)
}

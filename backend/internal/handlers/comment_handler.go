package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	repository "forum-project/backend/internal/repository/comments"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	comment := repository.Comment{}
	json.NewDecoder(r.Body).Decode(&comment)
	comment.Add()
	fmt.Println(comment)
}

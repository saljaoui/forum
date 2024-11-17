package handlers

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"

	comment "forum-project/backend/internal/repository/comments"
)


func Comment_handler(res http.ResponseWriter, req *http.Request) {
	myComment := comment.Comment{}
	json.NewDecoder(req.Body).Decode(&myComment)
	myComment.Add()
	fmt.Printf("comment with id  : %v  is created\n", myComment.ID)
	// json.NewDecoder(req.Body).Decode()
}

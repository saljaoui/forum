package handlers

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	comment "forum-project/backend/internal/repository/comments"
)

func convertToInteger(data  string) int{
	number,err := strconv.Atoi(data)
	if err != nil {
		return -1
	}
	return number
}

func Comment_handler(res http.ResponseWriter, req *http.Request) {
	userId := convertToInteger(req.FormValue("user_id"))
	content := req.FormValue("content")
	target,errTarget := strconv.Atoi(req.FormValue("target"))
	decode:=json.NewDecoder(req.Body)
	decode.DisallowUnknownFields()

	myComment := comment.NewComment()
	json.NewDecoder(req.Body).Decode(&myComment)
	myComment.Add()
	fmt.Printf("comment with id  : %v  is created\n", myComment.ID)
	// json.NewDecoder(req.Body).Decode()
}

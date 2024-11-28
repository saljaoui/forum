package handlers

import (
	//"encoding/json"
	"encoding/json"
	"strconv"

	//"fmt"
	comment "forum-project/backend/internal/repository/comments"
	"net/http"
)

func Comment_handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		id, err := strconv.Atoi(req.FormValue("target_id"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		comments := comment.GetAllCommentsbyTarget(id)
		if comments == nil {
			res.WriteHeader(http.StatusNotFound)
			return 
		}
		res.WriteHeader(http.StatusOK)
		encoder :=  json.NewEncoder(res)
		for _, c := range comments {
			encoder.Encode(c)	
		}
	} else if req.Method == "POST" {
		statusCode := addComment(req)
		if statusCode == http.StatusOK {
			res.Write([]byte("comment added succesfuly"))
			return 
		}
		if statusCode == http.StatusBadRequest {
			res.Write([]byte("comment Infos are wrongs!! "))
			return 
		}
		res.WriteHeader(statusCode)
	}
}

func addComment(req *http.Request) int {
	comment := comment.Comment{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		return http.StatusBadRequest
	}
	comment.Add()
	return http.StatusOK
}

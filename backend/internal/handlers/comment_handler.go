package handlers

import (
	//"encoding/json"
	"encoding/json"
	"strconv"

	//"fmt"
	comment "forum-project/backend/internal/repository/comments"
	"net/http"
)

func Handel_GetCommet(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		id, err := strconv.Atoi(req.FormValue("target_id"))
		if err != nil {
			JsoneResponse(res, "Status Bad Request", http.StatusBadRequest)
			return
		}
		comments := comment.GetAllCommentsbyTarget(id)
		if comments == nil {
			JsoneResponse(res, "Status Not Found", http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(res)
		for _, c := range comments {
			encoder.Encode(c)
		}
	} else {
		JsoneResponse(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
func Comment_handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		statusCode := addComment(req)
		if statusCode == http.StatusOK {
			JsoneResponse(res, "comment added succesfuly", http.StatusCreated)
			return
		}
		if statusCode == http.StatusBadRequest {
			JsoneResponse(res, "comment Infos are wrongs!! ", http.StatusBadRequest)
			return
		}
	} else {
		JsoneResponse(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
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

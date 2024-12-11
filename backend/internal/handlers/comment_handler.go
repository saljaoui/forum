package handlers

import (
	//"encoding/json"
	"encoding/json"
	"net/http"
	"strconv"

	//"fmt"
	comment "forum-project/backend/internal/repository/comments"
)

func Handel_GetCommet(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		HandleError(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(req.FormValue("target_id"))
	if err != nil {
		HandleError(res, "Status Bad Request", http.StatusBadRequest)
		return
	}

	comments := comment.GetAllCommentsbyTarget(id)
	if comments == nil {
		HandleError(res, "Status Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(res).Encode(comments)
}

func Handler_AddComment(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		HandleError(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	statusCode := addComment(req)
	if statusCode == http.StatusBadRequest {
		HandleError(res, "comment Infos are wrongs!! ", http.StatusBadRequest)
		return
	}

	if statusCode == http.StatusOK {
		JsoneResponse(res, "comment added succesfuly", http.StatusCreated)
		return
	}
}

func addComment(req *http.Request) int {
	var comment comment.Comment
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&comment)
	if err != nil {
		return http.StatusBadRequest
	}

	comment.Add()
	
	return http.StatusOK
}

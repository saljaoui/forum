package handlers

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
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
	// encoder := NewEncoderJsone(res)
	// for _, c := range comments {
	// 	err := encoder.Encode(c)
	// 	if err != nil {
	// 		JsoneResponse(res, "Error Encoding Comment", http.StatusInternalServerError)
	// 		return
	// 	}
	// }
	json.NewEncoder(res).Encode(comments)
	// JsoneResponse(res, comments, http.StatusOK)
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
//82a3abe1-39e9-47f5-bb1d-1ade395d4206//82a3abe1-39e9-47f5-bb1d-1ade395d4206

func addComment(req *http.Request) int {
	iduser := GetUserId(req)
	comment := comment.Comment{}
 	comment.User_Id = iduser
	if comment.User_Id == 0 {
		fmt.Println("error")
		return -1
	}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		return http.StatusBadRequest
	}
	comment.Add()
	return http.StatusOK
}

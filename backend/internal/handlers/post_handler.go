package handlers

import (
	"fmt"
	"net/http"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/posts"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	// if r.URL.Path != "/api/post" {
	// 	JsoneResponse(w, r, "Invalid path", http.StatusBadRequest)
	// 	return

	// }
	if r.Method != http.MethodPost {
		JsoneResponse(w, r, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id_user := GetUserId(r)
	post := posts.Post{}

	decode := DecodeJson(r)
	err := decode.Decode(&post)
	if err != nil {
		JsoneResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	if(checkdeblicat(post.Name_Category)){
		fmt.Println("Duplicate")
		JsoneResponse(w, r, "Duplicate category: The category already exists", http.StatusConflict)	
		return 
	}
	for _, n := range post.Name_Category {
		if !checkGategory(n) {
			JsoneResponse(w, r, "Your category is incorrect", http.StatusBadRequest)
			return
		}
	}
	post.User_Id = id_user
	post.CheckPostErr(w)
	id := post.Add()

	for _, name := range post.Name_Category {
		err := category.AddCategory(id, name)
		if err != nil {
			JsoneResponse(w, r, "Failed to add category", http.StatusBadRequest)
			return
		}
	}
	JsoneResponse(w, r, "create post Seccessfuly", http.StatusCreated)
}
func checkdeblicat(cat []string) bool{
	for i:=0;i<len(cat);i++{
		for j:=i+1;j<len(cat);j++{
			if cat[i]==cat[j]{
				return true
			}
		}
	}
	return false
}
func checkGategory(name string) bool {
	cate := []string{
		"General",
		"Technology",
		"Sports",
		"Entertainment",
		"Science",
		"Health",
		"Food",
		"Travel",
		"Fashion",
		"Art",
		"Music",
	}
	for _, v := range cate {
		if v == name {
			return true
		}
	}
	return false
}

package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	category "forum-project/backend/internal/repository/categories"
	"forum-project/backend/internal/repository/posts"

	"github.com/google/uuid"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
	if checkdeblicat(post.Name_Category) {
		JsoneResponse(w, r, "Duplicate category: The category already exists", http.StatusConflict)
		return
	}
	if len(post.Content) > 1000 {
		JsoneResponse(w, r, "Your content is long", http.StatusBadRequest)
		return
	}
	if post.Content == "" {
		JsoneResponse(w, r, "Your content is emty", http.StatusBadRequest)
		return
	}

	if post.Image_url != "" {
		post.Image_url, err = ParseImage(post.Image_url)
		if err != nil {
			JsoneResponse(w, r, err.Error(), http.StatusBadRequest)
			return
		}
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

func checkdeblicat(cat []string) bool {
	for i := 0; i < len(cat); i++ {
		for j := i + 1; j < len(cat); j++ {
			if cat[i] == cat[j] {
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

func ParseImage(image_url string) (string, error) {
	parts := strings.SplitN(image_url, ",", 2)
	if len(parts) != 2 || !strings.HasPrefix(parts[0], "data:image/") {
		return "", errors.New("invalid image URL")
	}
	mime := parts[0]
	var ext string
	switch {
	case strings.Contains(mime, "png"):
		ext = ".png"
	case strings.Contains(mime, "jpeg") || strings.Contains(mime, "jpg"):
		ext = ".jpeg"
	case strings.Contains(mime, "gif"):
		ext = ".gif"
	default:
		return "", errors.New("unsupported image type")
	}

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Printf("error decoding image: %v\n", err)
		return "", errors.New("something went wrong")
	}

	const maxSize = 20 * 1024 * 1024
	if len(data) > maxSize {
		return "", errors.New("image size exceeds 20MB")
	}
	imagePath := fmt.Sprintf("uploaded/%s%s", uuid.New().String(), ext)
	if err := os.MkdirAll("../../frontend/static/imgs/uploaded", 0755); err != nil {
		fmt.Printf("error creating directory: %v\n", err)
		return "", errors.New("something went wrong")
	}
	if err := os.WriteFile("../../frontend/static/imgs/"+imagePath, data, 0644); err != nil {
		fmt.Println(err.Error())
		return "", errors.New("failed to save image")
	}
	return imagePath, nil
}

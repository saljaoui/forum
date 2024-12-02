package posts

import (
	"encoding/json"
	"net/http"
	"time"

	"forum-project/backend/internal/database"
	"forum-project/backend/internal/repository/cards"
	like "forum-project/backend/internal/repository/likes"
)

type Post struct {
	ID            int      `json:"id"`
	User_Id       int      `json:"user_id"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Name_Category []string `json:"name"`
	CreatedAt     string   `json:"createdat"`
	Card_Id       int      `json:"card_id"`
}

type PostResponde struct {
	Card_Id   int
	Post_Id   int
	UserID    int
	FirstName string
	LastName  string
	Title     string
	Content   string
	Likes     int
	Dislikes  int
	CreatedAt time.Time
}

func (p *Post) Add() int {
	card := cards.NewCard(p.User_Id, p.Content)
	card.Add()
	if card.Id == -1 {
		return -1
	}
	p.Card_Id = card.Id
	id_posr := inserPost(p.Title, p.Card_Id)

	return int(id_posr)
}

func (p *Post) CheckPostErr(w http.ResponseWriter) {
	if p.Title == "" || p.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid input")
	}
}

func GetPosts(query string) []PostResponde {
	db := database.Config()
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var posts []PostResponde
	for rows.Next() {
		var post PostResponde
		err := rows.Scan(
			&post.Card_Id,
			&post.Post_Id,
			&post.UserID,
			&post.FirstName,
			&post.LastName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			return nil
		}
		likes, dislikes := like.GetLikes(post.Post_Id)
		post.Likes = likes
		post.Dislikes = dislikes
		posts = append(posts, post)
	}
	return posts
}

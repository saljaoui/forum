package like

import (
	"fmt"
	"strconv"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
)

func inserLike(user_id, card_id, is_liked int) (m messages.Messages) {
	if likeExists(user_id, card_id) {
		m.MessageError = "user already liked or desliked this"
		fmt.Println("user already liked or desliked this")
		return m
	}
	query := "INSERT INTO likes(user_id, card_id, is_like) VALUES(?,?,?);"
	_, err := database.Exec(query, user_id, card_id, is_liked)
	if err != nil {
		fmt.Println(err.Error())
	}
	m.MessageSucc = "is liked"
	return m
}

func deletLike(user_id, card_id int) {
	query := "DELETE FROM likes WHERE user_id=? AND card_id=?"
	_, err := database.Exec(query, user_id, card_id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetLikes(post_id int) (int, int) {
	querylike := `SELECT sum(is_like) FROM post p, likes l WHERE p.card_id = l.card_id AND l.is_like = 1 AND p.id = ` + strconv.Itoa(post_id)
	like := 0
	err := database.SelectOneRow(querylike).Scan(&like)
	if err != nil {
		like = 0
	}
	querydislike := `SELECT sum(is_like) FROM post p, likes l WHERE p.card_id = l.card_id AND l.is_like = -1 AND p.id = ` + strconv.Itoa(post_id)
	dislike := 0
	err = database.SelectOneRow(querydislike).Scan(&dislike)
	if err != nil {
		dislike = 0
	}
	return like, dislike * -1
}

func likeExists(user_id, card_id int) bool {
	var exists bool
	query := "SELECT EXISTS (select is_like from likes where user_id = ? AND card_id = ?)"
	err := database.SelectOneRow(query, user_id, card_id).Scan(&exists)
	if err != nil {
		fmt.Println("Error exist Like", err)
	}
	return exists
}

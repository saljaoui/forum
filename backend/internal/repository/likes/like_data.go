package like

import (
	"fmt"
	"strconv"

	"forum-project/backend/internal/database"
)

func inserLike(user_id, card_id, is_liked int) {
	query := "INSERT INTO likes(user_id, card_id, is_like) VALUES(?,?,?);"
	_, err := database.Exec(query, user_id, card_id, is_liked)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func deletLike( user_id,card_id int) {
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

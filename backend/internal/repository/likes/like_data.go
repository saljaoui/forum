package like

import (
	"fmt"

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

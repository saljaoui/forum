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

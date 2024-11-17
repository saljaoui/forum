package posts

import (
	"forum-project/backend/internal/database"
)

func inserPost(title string, card_id int) int {
	query := "INSERT INTO post(title, card_id) VALUES(?,?);"
	return database.Exec(query, title, card_id)
}

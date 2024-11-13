package repository

import (
	"fmt"

	"forum-project/backend/internal/database"

	"github.com/gofrs/uuid/v5"
)

func emailExists(email string) bool {
	var exists bool
	query := "SELECT EXISTS (select email from user where email=?)"
	database.SelectOneRow(query, email, &exists)
	return exists
}

func updateuuidUser(uudi uuid.UUID, userId int64) {
	db := database.Config()
	_, err := db.Exec("UPDATE user SET UUID=? WHERE id=?", uudi, userId)
	if err != nil {
		fmt.Println("Error To Update User uuid")
	}
}


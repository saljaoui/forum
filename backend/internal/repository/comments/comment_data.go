package comment

import "forum-project/backend/internal/database"


func insertComment(card_id,target_id int) int {
    query := "INSERT INTO comment(card_id,target_id) VALUES(?,?);"
    return database.Exec(query,card_id,target_id)
}